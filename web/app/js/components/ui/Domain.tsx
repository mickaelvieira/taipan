import React from "react";
import { Bookmark } from "../../types/bookmark";
import { Document } from "../../types/document";
import { Subscription } from "../../types/subscription";
import { ExternalLink } from "./Link";

export interface DomainProps {
  item: Bookmark | Document | Subscription;
  className?: string;
}

export default React.memo(function Domain({
  item,
  className
}: DomainProps): JSX.Element {
  const url = new URL(item.url);
  return (
    <ExternalLink
      href={item.url}
      title={item.title ? item.title : item.url}
      className={className}
    >
      {url.host}
    </ExternalLink>
  );
});
