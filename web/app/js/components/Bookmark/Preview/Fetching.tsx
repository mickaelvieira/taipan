import { wrap } from "dom-element-wrapper";
import Loader from "components/Common/Loader";

const Fetching = () => {
  const wrappers = [
    wrap("div", { className: "col bookmark-col" }),
    wrap("div", { className: "col bookmark-loader spinner-red" }).append(
      Loader()
    ),
    wrap("div", { className: "col bookmark-col" })
  ];

  return wrappers.map(element => element.unwrap());
};

export default Fetching;
