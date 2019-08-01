import React, { PropsWithChildren } from "react";
import Link, { LinkProps } from "@material-ui/core/Link";

export default function ExternalLink({
  children,
  ...props
}: PropsWithChildren<LinkProps>): JSX.Element {
  return (
    <Link target="_blank" rel="noopener" {...props}>
      {children}
    </Link>
  );
}
