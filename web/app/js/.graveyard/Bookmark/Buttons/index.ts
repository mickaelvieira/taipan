import { wrap } from "dom-element-wrapper";
import ButtonClose from "./ButtonClose";
import ButtonLink from "./ButtonLink";

const Buttons = ({ item }) =>
  wrap("div", {
    className: "bookmark-buttons-container"
  }).append(ButtonLink(item), ButtonClose());

export default Buttons;
