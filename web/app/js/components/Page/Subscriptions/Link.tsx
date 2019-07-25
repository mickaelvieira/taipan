import React from "react";
import Link from "@material-ui/core/Link";
import { Subscription } from "../../../types/subscription";
import { Source } from "../../../types/syndication";
import { getDomain } from "../../../helpers/syndication";

interface Props {
  item: Subscription | Source;
}

export default React.memo(function SubscriptionLink({
  item
}: Props): JSX.Element {
  const url = getDomain(item);

  return (
    <Link
      underline="none"
      href={`${url.protocol}//${url.host}`}
      title={item.title ? item.title : item.url}
      target="_blank"
      rel="noopener"
    >
      {`${url.protocol}//${url.host}`}
    </Link>
  );
});
