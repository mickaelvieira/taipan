import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import { Bookmark } from "../../../types/bookmark";
import FeedItem from "./Item";

const useStyles = makeStyles({
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

export default function List({ hasResults, bookmarks }: Props) {
  const classes = useStyles();
  return !hasResults ? null : (
    <div className={classes.container}>
      {bookmarks.map((bookmark: Bookmark, index) => (
        <FeedItem bookmark={bookmark} index={index} key={bookmark.id} />
      ))}
    </div>
  );
}
