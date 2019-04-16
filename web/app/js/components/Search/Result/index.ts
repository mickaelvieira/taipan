import { wrap } from "dom-element-wrapper";
import withHistory from "components/HoC/withHistory";
import { fetchBookmark, showBookmark } from "../../../actions";
import withDispatch from "components/HoC/withDispatch";

const Result = ({ dispatch, history, id, url, title, description }) => {
  const link = wrap("a", {
    href: "#"
  })
    .append(
      wrap("div", {
        className: "search-bookmark-result-title",
        innerHTML: title ? title : url.hostname
      }),
      wrap("div", {
        className: "search-bookmark-result-description"
        // innerHTML: description
      }).append(...description)
    )
    .unwrap();

  link.addEventListener("click", event => {
    event.preventDefault();
    history.pushState(null, title, "/bookmark/" + id);
    dispatch(showBookmark());
    dispatch(fetchBookmark(id));
  });

  return wrap("li", { className: "search-bookmark-result" }).append(link);
};

export default withHistory(withDispatch(Result));
