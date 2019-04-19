import { wrap } from "dom-element-wrapper";
import Loader from "components/Common/Loader";

const Fetching = () =>
  wrap("li", {
    className: "search-bookmark-fetching spinner-red"
  }).append(Loader());

export default Fetching;
