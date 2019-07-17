import React from "react";
import { Bookmark } from "../../../types/bookmark";
import Item from "./Item";
import { ListProps } from "../../ui/Feed/Feed";
import FeedItem from "../../ui/Feed/Item/Item";

export default function BookmarkList({
  results,
  updater
}: ListProps): JSX.Element {
  return (
    <>
      {results.length === 0} <div>You have no favorites yet :(</div>}
      {results.map(result => (
        <FeedItem item={result} updater={updater} key={result.id}>
          {({ remove }) => (
            <Item bookmark={result as Bookmark} remove={remove} />
          )}
        </FeedItem>
      ))}
    </>
  );
}
