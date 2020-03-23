import React from "react";
import { Document } from "../../../../types/document";
import Item from "../../../ui/Feed/Item/Document";
import { ListProps } from "../../../ui/Feed/Feed";
import FeedItem from "../../../ui/Feed/Item/Wrapper";
import NoResults from "../NoResults";
import { SearchProps } from "../";

export default React.memo(function DocumentList({
  results,
  terms,
  type,
}: SearchProps & ListProps): JSX.Element {
  return (
    <>
      {results.length === 0 && <NoResults terms={terms} type={type} />}
      {results.map((result) => (
        <FeedItem key={result.id}>
          <Item document={result as Document} />
        </FeedItem>
      ))}
    </>
  );
});
