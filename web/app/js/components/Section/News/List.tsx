import React from "react";
import { Document } from "../../../types/document";
import Item from "./Item";

interface Props {
  results: Document[];
}

export default function List({ results }: Props) {
  return results.map((result, index) => (
    <Item document={result} index={index} key={result.id} />
  ));
}
