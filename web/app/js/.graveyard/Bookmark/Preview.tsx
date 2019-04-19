import { wrap } from "dom-element-wrapper";
import Info from "./Info/index";
import Fetching from "./Preview/Fetching";
import Content from "./Preview/Content";
import Buttons from "./Buttons/index";

const Preview = ({ html, data, history, visible }) => {
  if (!data || !visible) {
    return Fetching();
  }

  const wrappers = [
    wrap("div", { className: "col bookmark-col" }),
    wrap("div", {
      className: "col bookmark-container"
    })
      .append(Buttons({ item: data }))
      .append(Content({ html }))
      .append(Info({ data, history })),
    wrap("div", { className: "col bookmark-col" })
  ];

  return wrappers.map(wrapper => wrapper.unwrap());
};

export default Preview;
