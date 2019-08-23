import React from "react";
import { Bookmark } from "../../../../types/bookmark";
import Item from "../../../ui/Feed/Item/Bookmark";
import { ListProps } from "../../../ui/Feed/Feed";
import FeedItem from "../../../ui/Feed/Item/Wrapper";
import NoResults from "../NoResults";
import { SearchProps } from "../";

export default React.memo(function BookmarkList({
  results,
  terms,
  type
}: SearchProps & ListProps): JSX.Element {
  return (
    <>
      {results.length === 0 && <NoResults terms={terms} type={type} />}
      {results.map(result => (
        <FeedItem key={result.id}>
          <Item bookmark={result as Bookmark} />
        </FeedItem>
      ))}
    </>
  );
});
