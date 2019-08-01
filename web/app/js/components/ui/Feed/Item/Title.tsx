import React from "react";
import Typography from "@material-ui/core/Typography";
import ExternalLink from "../../../ui/Link/External";
import { Bookmark } from "../../../../types/bookmark";
import { Document } from "../../../../types/document";

interface Props {
  item: Document | Bookmark;
}

export default function ItemTitle({ item }: Props): JSX.Element {
  return (
    <ExternalLink href={item.url} title={item.title}>
      <Typography variant="h6" component="h6">
        {item.title}
      </Typography>
    </ExternalLink>
  );
}
