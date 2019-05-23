import React from "react";
import { withStyles, WithStyles, createStyles } from "@material-ui/core/styles";
import { Bookmark } from "../../../types/bookmark";
import FeedItem from "./Item";

const styles = () =>
  createStyles({
    container: {
      display: "flex",
      flexDirection: "column",
      margin: 12
    }
  });

export interface Props {
  hasResults: boolean;
  bookmarks: Bookmark[];
}

export default withStyles(styles)(function List({
  hasResults,
  bookmarks,
  classes
}: Props & WithStyles<typeof styles>) {
  return !hasResults ? null : (
    <div className={classes.container}>
      {bookmarks.map((bookmark: Bookmark, index) => (
        <FeedItem bookmark={bookmark} index={index} key={bookmark.id} />
      ))}
    </div>
  );
});
