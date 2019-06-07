import React from "react";
import { Document } from "../../../types/document";
import Item from "./Item";
import { ListProps } from "../../ui/Feed/Feed";

export default function DocumentList({ results, query }: ListProps) {
  return (
    <>
      {results.map((result, index) => (
        <Item
          document={result as Document}
          index={index}
          key={result.id}
          query={query}
        />
      ))}
    </>
  );
}
