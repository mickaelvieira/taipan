import { wrap } from "dom-element-wrapper";

const Source = ({ title, url }) =>
  wrap("div").append(
    wrap("strong").append("Source: "),
    wrap("a", {
      title,
      href: url,
      target: "_blank"
      // className: "red"
    }).append(new URL(url).hostname)
  );

export default Source;
