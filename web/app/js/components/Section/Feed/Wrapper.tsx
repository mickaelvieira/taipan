import React, { useEffect } from "react";
import Loader from "../../ui/Loader";
import List, { Props as ListProps } from "./List";
import { variables, Variables, Data } from "../../apollo/Query/Feeds";
import useWindowBottom from "../../../hooks/window-bottom";

interface Props extends ListProps {
  isLoading: boolean;
  queryKey: string;
  fetchMore: (options: {
    variables: Variables;
    updateQuery: (
      previousQueryResult: Data,
      options: {
        fetchMoreResult?: Data;
        variables?: Variables;
      }
    ) => Data;
  }) => Promise<Data>;
}

export default function FeedWrapper({
  isLoading,
  queryKey,
  fetchMore,
  hasResults,
  bookmarks
}: Props) {
  const isAtTheBottom = useWindowBottom();
  const offset = bookmarks.length;

  // useEffect(() => {
  // console.log("fetch more");
  // console.log(`Bottom ${isAtTheBottom}`);
  // console.log(`Loading ${isLoading}`);
  // console.log(`Offset ${offset}`);

  //   if (isAtTheBottom && !isLoading && offset > 0) {
  //     fetchMore({
  //       variables: {
  //         ...variables,
  //         offset: offset
  //       },
  //       updateQuery: (prev, { fetchMoreResult }) => {
  //         console.log("update query");
  //         console.log(prev);
  //         console.log(fetchMoreResult);

  //         if (!fetchMoreResult) {
  //           return prev;
  //         }

  //         return {
  //           [queryKey]: {
  //             ...fetchMoreResult[queryKey],
  //             results: [
  //               ...prev[queryKey].results,
  //               ...fetchMoreResult[queryKey].results
  //             ]
  //           }
  //         };
  //       }
  //     });
  //   }
  // }, [offset, fetchMore, isAtTheBottom, isLoading]);

  return (
    <>
      {isLoading && <Loader />}
      {!isLoading && hasResults && (
        <List hasResults={hasResults} bookmarks={bookmarks} />
      )}
    </>
  );
}
