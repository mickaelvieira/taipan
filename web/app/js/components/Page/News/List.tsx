import React from "react";
import { Document } from "../../../types/document";
import Item from "./Item";
import { ListProps } from "../../ui/Feed/Feed";
import Latest from "./Latest";
import FeedItem from "../../ui/Feed/Item/Item";
import EmptyFeed from "../../ui/Feed/Empty";

export default function DocumentList({
  results,
  firstId,
  lastId,
  updater
}: ListProps): JSX.Element {
  return (
    <>
      <Latest firstId={firstId} lastId={lastId} />
      {results.length === 0 && <EmptyFeed message="No news today" />}
      {results.map(result => (
        <FeedItem item={result} updater={updater} key={result.id}>
          {({ remove }) => (
            <Item document={result as Document} remove={remove} />
          )}
        </FeedItem>
      ))}
    </>
  );
}
