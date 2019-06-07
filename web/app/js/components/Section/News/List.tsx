import React from "react";
import { Document } from "../../../types/document";
import Item from "./Item";
import { ListProps } from "../../ui/Feed/Feed";

export default function DocumentList({ results, query }: ListProps) {
  return (
    <>
      {results.map(result => (
        <Item document={result as Document} key={result.id} query={query} />
      ))}
    </>
  );
}
