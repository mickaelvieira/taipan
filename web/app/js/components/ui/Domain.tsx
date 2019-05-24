import React from "react";
import Link from "@material-ui/core/Link";
import { Bookmark } from "../../types/bookmark";
import { Document } from "../../types/document";
import { withStyles, WithStyles, createStyles } from "@material-ui/core/styles";

const styles = () =>
  createStyles({
    link: {
      padding: 12
    }
  });

interface Props extends WithStyles<typeof styles> {
  item: Bookmark | Document;
}

export default withStyles(styles)(
  React.memo(function Domain({ classes, item }: Props) {
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
  })
);
