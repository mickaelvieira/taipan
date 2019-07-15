import React, { useState } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Fade from "@material-ui/core/Fade";
import Card from "@material-ui/core/Card";
import { FeedUpdater } from "../../../apollo/helpers/feed";
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

interface Props {
  item: FeedItem;
  updater: FeedUpdater;
  children: (props: RenderProps) => JSX.Element;
}

interface RenderProps {
  remove: () => void;
}

export default function Item({ children, updater, item }: Props): JSX.Element {
  const classes = useStyles();
  const [visible, setIsVisible] = useState(true);

  return (
    <Fade
      in={visible}
      unmountOnExit
      timeout={{
        enter: 500,
        exit: 400
      }}
      onExited={() => {
        updater(item, "Remove");
      }}
    >
      <Card className={classes.card}>
        {children({
          remove: () => {
            setIsVisible(false);
          }
        })}
      </Card>
    </Fade>
  );
}
