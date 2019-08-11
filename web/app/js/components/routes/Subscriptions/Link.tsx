import React from "react";
import { ExternalLink } from "../../ui/Link";
import { Subscription } from "../../../types/subscription";
import { Source } from "../../../types/syndication";

interface Props {
  item: Subscription | Source;
}

export default React.memo(function SubscriptionLink({
  item
}: Props): JSX.Element {
  return (
    <ExternalLink href={item.url} title={item.title ? item.title : item.url}>
      {item.url}
    </ExternalLink>
  );
});
