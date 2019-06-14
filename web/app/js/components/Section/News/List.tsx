import React, { useContext, useEffect } from "react";
import { Document } from "../../../types/document";
import Item from "./Item";
import { ListProps } from "../../ui/Feed/Feed";

import { NewsContext } from "../../context";

export default function DocumentList({ results, query }: ListProps) {
  const setToId = useContext(NewsContext);
  console.log("render news");
  console.log(results);
  let firstId = "";
  if (results.length > 0) {
    firstId = results[0].id;
  }

  useEffect(() => {
    setToId(firstId);
  }, [firstId, setToId]);

  return (
    <>
      {results.map(result => (
        <Item document={result as Document} key={result.id} query={query} />
      ))}
    </>
  );
}
