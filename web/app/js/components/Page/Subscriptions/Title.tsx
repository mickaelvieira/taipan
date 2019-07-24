import React from "react";
import Link from "@material-ui/core/Link";
import Typography from "@material-ui/core/Typography";
import { Subscription } from "../../../types/subscription";

interface Props {
  item: Subscription;
}

export default function SubscriptionTitle({ item }: Props): JSX.Element {
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
