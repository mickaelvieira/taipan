import React from "react";
import { Bookmark } from "../../../types/bookmark";
import Item from "./Item";
import { ListProps } from "../../ui/Feed/Feed";

export default function BookmarkList({ results }: ListProps): JSX.Element {
  return (
    <>
      {results.map(result => (
        <Item bookmark={result as Bookmark} key={result.id} />
      ))}
    </>
  );
}
