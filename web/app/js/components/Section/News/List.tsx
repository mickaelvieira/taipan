import React from "react";
import PropTypes from "prop-types";
import { Document } from "../../../types/document";
import Item from "./Item";

interface Props {
  results: Document[];
  query: PropTypes.Validator<object>;
}

export default function DocumentList({ results, query }: Props) {
  return results.map((result, index) => (
    <Item document={result} index={index} key={result.id} query={query} />
  ));
}
