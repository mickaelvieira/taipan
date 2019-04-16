import { wrap } from "dom-element-wrapper";

const NoResult = () =>
  wrap("li", { className: "search-bookmark-no-results" }).append("No results");

export default NoResult;
