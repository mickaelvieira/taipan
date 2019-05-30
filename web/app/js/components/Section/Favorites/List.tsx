import React from "react";
import { Bookmark } from "../../../types/bookmark";
import Item from "./Item";

interface Props {
  results: Bookmark[];
}

export default function List({ results }: Props) {
  return results.map((result, index) => (
    <Item bookmark={result} index={index} key={result.id} />
  ));
}
