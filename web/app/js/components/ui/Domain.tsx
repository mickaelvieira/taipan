import React from "react";
import Link from "@material-ui/core/Link";
import { Bookmark } from "../../types/bookmark";
import { Document } from "../../types/document";
import { Source } from "../../types/syndication";
import { makeStyles } from "@material-ui/core/styles";

const useStyles = makeStyles({
  link: {
    padding: 12
  }
});

interface Props {
  item: Bookmark | Document | Source;
}

export default React.memo(function Domain({ item }: Props): JSX.Element {
  const classes = useStyles();
  const url = new URL(item.url);
  return (
    <Link
      underline="none"
      href={item.url}
      title={item.title ? item.title : item.url}
      target="_blank"
      rel="noopener"
      className={classes.link}
    >
      {url.host}
    </Link>
  );
});
