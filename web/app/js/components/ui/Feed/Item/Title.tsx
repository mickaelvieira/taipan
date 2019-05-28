import React from "react";
import Link from "@material-ui/core/Link";
import Typography from "@material-ui/core/Typography";
import { Bookmark } from "../../../../types/bookmark";
import { Document } from "../../../../types/document";

interface Props {
  item: Document | Bookmark;
}

export default function ItemTitle({ item }: Props) {
  return (
    <Link
      underline="none"
      href={item.url}
      title={item.title}
      target="_blank"
      rel="noopener"
    >
      <Typography variant="h6" component="h6">
        {item.title}
      </Typography>
    </Link>
  );
}
