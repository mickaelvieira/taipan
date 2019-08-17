import React from "react";
import { ExternalLink } from "../../ui/Link";
import { Subscription } from "../../../types/subscription";
import { Source } from "../../../types/syndication";
import { getDomain } from "../../../helpers/syndication";

interface Props {
  item: Subscription | Source;
}

export default React.memo(function SubscriptionDomain({
  item
}: Props): JSX.Element {
  const url = getDomain(item);

  return (
    <ExternalLink
      href={`${url.protocol}//${url.host}`}
      title={item.title ? item.title : `${item.url}`}
    >
      {`${url.protocol}//${url.host}`}
    </ExternalLink>
  );
});
