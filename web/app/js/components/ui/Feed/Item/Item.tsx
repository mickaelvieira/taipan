import React, { ReactNode, useState } from "react";
import PropTypes from "prop-types";
import { ApolloConsumer } from "react-apollo";
import { makeStyles } from "@material-ui/core/styles";
import Fade from "@material-ui/core/Fade";
import Card from "@material-ui/core/Card";
import { getDataKey } from "../../../apollo/Query/Feed";
import { feedResultsAction } from "../../../apollo/helpers/feed";
import { FeedItem } from "../../../../types/feed";

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
  item: FeedItem;
  children: (props: RenderProps) => ReactNode;
  query: PropTypes.Validator<object>;
}

export default function Item({ children, query, item }: Props) {
  const classes = useStyles();
  const [visible, setIsVisible] = useState(true);

  return (
    <ApolloConsumer>
      {client => (
        <Fade
          in={visible}
          unmountOnExit
          timeout={{
            enter: 500,
            exit: 400
          }}
          onExited={() => {
            try {
              const data = client.readQuery({ query });
              const updateResults = feedResultsAction["Remove"];
              if (data) {
                const key = getDataKey(data);
                if (key) {
                  const result = updateResults(data[key], item);
                  client.writeQuery({
                    query,
                    data: { [key]: result }
                  });
                }
              }
            } catch (e) {
              console.warn(e);
            }
          }}
        >
          <Card className={classes.card}>
            {children({
              remove: () => {
                // setIsVisible(false)
              }
            })}
          </Card>
        </Fade>
      )}
    </ApolloConsumer>
  );
}
