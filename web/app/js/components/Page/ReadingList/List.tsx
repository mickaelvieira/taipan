import React from "react";
import { Bookmark } from "../../../types/bookmark";
import Item from "./Item";
import { ListProps } from "../../ui/Feed/Feed";
import FeedItem from "../../ui/Feed/Item/Item";
import EmptyFeed from "../../ui/Feed/Empty";

export default function BookmarkList({ results }: ListProps): JSX.Element {
  return (
    <>
      {results.length === 0 && (
        <EmptyFeed message="Well done your reading list is empty \o/" />
      )}
      {results.map(result => (
        <FeedItem key={result.id}>
          {({ remove }) => (
            <Item bookmark={result as Bookmark} remove={remove} />
          )}
        </FeedItem>
      ))}
    </>
  );
}
