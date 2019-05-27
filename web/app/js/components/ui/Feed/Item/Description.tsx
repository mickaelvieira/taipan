import React from "react";
import Typography from "@material-ui/core/Typography";
import { Bookmark } from "../../../../types/bookmark";
import { Document } from "../../../../types/document";
import { truncate } from "../../../../helpers/string";

interface Props {
  item: Document | Bookmark;
}

export default function ItemDescription({ item }: Props) {
  return !item.description ? null : (
    <Typography gutterBottom>{truncate(item.description)}</Typography>
  );
}
