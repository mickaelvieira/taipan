import React from "react";
import { Bookmark } from "../../../types/bookmark";
import Item from "./Item";
import { ListProps } from "../../ui/Feed/Feed";
import FeedItem from "../../ui/Feed/Item/Item";
import EmptyFeed from "../../ui/Feed/Empty";
import Emoji from "../../ui/Emoji";

export default React.memo(function BookmarkList({
  results
}: ListProps): JSX.Element {
  return (
    <>
      {results.length === 0 && (
        <EmptyFeed>
          <span>
            Well done your reading list is empty
            <Emoji emoji=":i_love_you_hand_sign:" />
          </span>
        </EmptyFeed>
      )}
      {results.map(result => (
        <FeedItem key={result.id}>
          <Item bookmark={result as Bookmark} />
        </FeedItem>
      ))}
    </>
  );
});
