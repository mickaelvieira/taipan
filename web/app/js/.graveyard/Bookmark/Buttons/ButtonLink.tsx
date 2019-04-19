import { wrap } from "dom-element-wrapper";

const ButtonLink = ({ url }) =>
  wrap("a", {
    href: url,
    target: "_blank",
    className: "link-source"
  }).append(wrap("i", { className: "fa fa-link" }));

export default ButtonLink;
