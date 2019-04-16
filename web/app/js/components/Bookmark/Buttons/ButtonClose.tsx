import { wrap } from "dom-element-wrapper";
import { hideBookmark } from "../../../actions";
import withDispatch from "components/HoC/withDispatch";

const ButtonClose = ({ dispatch }) =>
  wrap("button", { className: "btn-link btn-close" })
    .append(wrap("i", { className: "fa fa-close" }))
    .addEventListener("click", () => {
      window.history.pushState(null, null, "/");
      dispatch(hideBookmark());
    });

export default withDispatch(ButtonClose);
