import React from "react";
import Link from "@material-ui/core/Link";
import Typography from "@material-ui/core/Typography";
import { Source } from "../../../types/syndication";

interface Props {
  item: Source;
}

export default function SourceTitle({ item }: Props): JSX.Element {
  return (
    <Link
      underline="none"
      href={item.url}
      title={item.title}
      target="_blank"
      rel="noopener"
    >
      <Typography component="span" style={{ padding: "12px" }}>
        {item.title}
      </Typography>
    </Link>
  );
}
