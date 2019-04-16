import { wrap } from "dom-element-wrapper";
import moment from "moment";
import {
  getLastEntry,
  getLastFailure,
  getLastSuccess
} from "lib/history/bookmarks";
import LabelledDate from "../Preview/LabelledDate";
import Controller from "./Controller";
import Source from "../Preview/Source";

const Info = ({ accessed_at, history }) => {
  const lastEntry = getLastEntry(history);
  const lastFailure = getLastFailure(history);
  const lastSuccess = getLastSuccess(history);
  const lastAccessed = accessed_at;
  // const url = new URL(data.url);

  const last = [
    ["Access", lastAccessed],
    ["Entry", lastEntry],
    ["Failure", lastFailure],
    ["Success", lastSuccess]
  ]
    .map(([type, entry]) => {
      return entry
        ? LabelledDate({
            label: `Last ${type}: `,
            date: moment(entry.created_at).fromNow()
          })
        : false;
    })
    .filter(entry => !!entry);

  const info = wrap("div", { className: "bookmark-info-container" })
    .append(
      wrap("button", { className: "btn-link btn-close" }).append(
        wrap("i", {
          className: "fa fa-chevron-up"
        })
      )
    )
    .append(
      Source(data),
      LabelledDate({
        label: "Added: ",
        date: moment(data.added_at).fromNow()
      })
    )
    .append(...last)
    .unwrap();

  new Controller(info);

  return info;
};

export default Info;
