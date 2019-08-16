import React from "react";
import { Bookmark } from "../../../types/bookmark";
import Item from "./Item";
import { ListProps } from "../../ui/Feed/Feed";
import FeedItem from "../../ui/Feed/Item/Item";
import EmptyFeed from "../../ui/Feed/Empty";

export default React.memo(function BookmarkList({
  results
}: ListProps): JSX.Element {
  return (
    <>
      {results.length === 0 && (
        <EmptyFeed message="You have no favorites yet :(" />
      )}
      {results.map(result => (
        <FeedItem key={result.id}>
          <Item bookmark={result as Bookmark} />
        </FeedItem>
      ))}
    </>
  );
});
