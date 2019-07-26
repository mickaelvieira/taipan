import React from "react";
import Link from "@material-ui/core/Link";
import { Subscription } from "../../../types/subscription";
import { Source } from "../../../types/syndication";

interface Props {
  item: Subscription | Source;
}

export default React.memo(function SubscriptionLink({
  item
}: Props): JSX.Element {
  return (
    <Link
      underline="none"
      href={item.url}
      title={item.title ? item.title : item.url}
      target="_blank"
      rel="noopener"
    >
      {item.url}
    </Link>
  );
});
