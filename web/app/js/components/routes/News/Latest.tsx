import React, { useContext } from "react";
import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import { Document } from "../../../types/document";
import { FeedsCacheContext } from "../../context";
import usePolling from "./usePolling";

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    container: {
      position: "fixed",
      backgroundColor: "#fff",
      alignSelf: "center",
    },
    button: {
      margin: theme.spacing(1),
    },
  })
);

interface Props {
  firstId?: string;
  lastId?: string;
}

export default function Latest({ firstId = "" }: Props): JSX.Element | null {
  const classes = useStyles();
  const updater = useContext(FeedsCacheContext);
  const [documents, flush] = usePolling(firstId);

  return documents.length === 0 ? null : (
    <div className={classes.container}>
      <Button
        className={classes.button}
        onClick={() => {
          if (!updater) {
            console.warn("Feed updater is not available");
            return;
          }
          // append documents to the feed
          updater.appendLatest(documents as Document[]);
          // reset the queue
          flush();
          window.scroll(0, 0);
        }}
      >
        See {documents.length} latest news
      </Button>
    </div>
  );
}
