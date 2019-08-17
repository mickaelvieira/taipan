import React from "react";
import { NavLink, LinkProps } from "react-router-dom";

const RouterLink = React.forwardRef<HTMLAnchorElement, LinkProps>(
  (props, ref) => (
    <NavLink exact innerRef={ref as React.Ref<HTMLAnchorElement>} {...props} />
  )
);

export default RouterLink;
