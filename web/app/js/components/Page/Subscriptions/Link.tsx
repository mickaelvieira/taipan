import React from "react";
import Link from "@material-ui/core/Link";
import { Subscription } from "../../../types/subscription";
import { getDomain } from "../../../helpers/syndication";

interface Props {
  subscription: Subscription;
}

export default React.memo(function SubscriptionLink({
  subscription
}: Props): JSX.Element {
  const url = getDomain(subscription);

  return (
    <Link
      underline="none"
      href={`${url.protocol}//${url.host}`}
      title={subscription.title ? subscription.title : subscription.url}
      target="_blank"
      rel="noopener"
    >
      {`${url.protocol}//${url.host}`}
    </Link>
  );
});
