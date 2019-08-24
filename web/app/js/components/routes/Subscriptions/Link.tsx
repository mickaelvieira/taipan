import React from "react";
import { ExternalLink } from "../../ui/Link";
import { Subscription } from "../../../types/subscription";

interface Props {
  item: Subscription;
}

export default React.memo(function SubscriptionLink({
  item
}: Props): JSX.Element {
  return (
    <ExternalLink
      href={`${item.url}`}
      title={item.title ? item.title : `${item.url}`}
    >
      {`${item.url}`}
    </ExternalLink>
  );
});
