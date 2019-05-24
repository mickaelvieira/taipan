import React from "react";
import Link from "@material-ui/core/Link";
import { Bookmark } from "../../types/bookmark";
import { Document } from "../../types/document";
import { makeStyles } from "@material-ui/core/styles";

const useStyles = makeStyles({
  link: {
    padding: 12
  }
});

interface Props {
  item: Bookmark | Document;
}

export default React.memo(function Domain({ item }: Props) {
  const classes = useStyles();
  const url = new URL(item.url);
  return (
    <Link
      underline="none"
      href={item.url}
      title={item.title}
      target="_blank"
      rel="noopener"
      className={classes.link}
    >
      {url.host}
    </Link>
  );
});
