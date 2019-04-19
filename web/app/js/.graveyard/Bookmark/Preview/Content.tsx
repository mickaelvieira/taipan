import { wrap } from "dom-element-wrapper";

const Content = ({ html }) =>
  wrap("hgroup", {
    className: "bookmark-content"
  }).insertAdjacentHTML("afterbegin", html ? html : "");

export default Content;
