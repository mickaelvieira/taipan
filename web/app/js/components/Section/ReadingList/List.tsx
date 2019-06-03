import React from "react";
import PropTypes from "prop-types";
import { Bookmark } from "../../../types/bookmark";
import Item from "./Item";

interface Props {
  results: Bookmark[];
  query: PropTypes.Validator<object>;
}

export default function BookmarkList({ results, query }: Props) {
  return results.map((result, index) => (
    <Item bookmark={result} index={index} key={result.id} query={query} />
  ));
}
