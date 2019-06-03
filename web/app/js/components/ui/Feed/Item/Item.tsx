import React, { ReactNode, useState } from "react";
import PropTypes from "prop-types";
import { ApolloConsumer } from "react-apollo";
import { makeStyles } from "@material-ui/core/styles";
import Fade from "@material-ui/core/Fade";
import Card from "@material-ui/core/Card";
import { getDataKey } from "../../../apollo/Query/Feed";

const useStyles = makeStyles(({ breakpoints }) => ({
  card: {
    marginBottom: 24,
    display: "flex",
    flexDirection: "column",
    borderRadius: 0,
    [breakpoints.up("sm")]: {
      borderRadius: 4
    }
  }
}));

interface RenderProps {
  remove: () => void;
}

interface Props {
  children: (props: RenderProps) => ReactNode;
  query: PropTypes.Validator<object>;
}

export default function Item({ children, query }: Props) {
  const classes = useStyles();
  const [visible, setIsVisible] = useState(true);

  return (
    <ApolloConsumer>
      {client => (
        <Fade
          in={visible}
          unmountOnExit
          timeout={{
            enter: 1000,
            exit: 500
          }}
          onExited={() => {
            const data = client.readQuery({ query });
            if (data) {
              const key = getDataKey(data);
              if (key) {
                client.query({
                  query,
                  variables: {
                    offset: 0,
                    limit: data[key].results.length - 1
                  },
                  fetchPolicy: "network-only"
                });
              }
            }
          }}
        >
          <Card className={classes.card}>
            {children({
              remove: () => setIsVisible(false)
            })}
          </Card>
        </Fade>
      )}
    </ApolloConsumer>
  );
}
