import React from "react";
import Link from "@material-ui/core/Link";
import { Bookmark } from "../../types/bookmark";
import { Document } from "../../types/document";
import { Subscription } from "../../types/subscription";

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
    <Link
      underline="none"
      href={item.url}
      title={item.title ? item.title : item.url}
      target="_blank"
      rel="noopener"
      className={className}
    >
      {url.host}
    </Link>
  );
});
