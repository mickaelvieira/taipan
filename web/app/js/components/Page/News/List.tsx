import React from "react";
import { Document } from "../../../types/document";
import Item from "./Item";
import { ListProps } from "../../ui/Feed/Feed";
import Latest from "./Latest";

export default function DocumentList({
  results,
  firstId,
  lastId
}: ListProps): JSX.Element {
  return (
    <>
      <Latest firstId={firstId} lastId={lastId} />
      {results.map(result => (
        <Item document={result as Document} key={result.id} />
      ))}
    </>
  );
}
