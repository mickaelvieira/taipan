import React from "react";
import { Bookmark } from "../../../types/bookmark";
import Item from "./Item";
import { ListProps } from "../../ui/Feed/Feed";

export default function BookmarkList({ results, query }: ListProps) {
  return (
    <>
      {results.map((result, index) => (
        <Item
          bookmark={result as Bookmark}
          index={index}
          key={result.id}
          query={query}
        />
      ))}
    </>
  );
}
