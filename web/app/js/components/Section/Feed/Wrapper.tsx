import React, { useEffect } from "react";
// import { FetchMoreOptions, ApolloQueryResult } from "react-apollo";
import { hasReachedTheBottom } from "../../../helpers/window";
import Loader from "../../ui/Loader";
import List, { Props as ListProps } from "./List";
import {
  variables,
  query,
  Variables,
  Data
} from "../../apollo/Query/LatestBookmarks";
import useWindoBottom from "../../../hooks/window-bottom";

// const styles = () =>
//   createStyles({
//     scrolling: {
//       pointerEvents: "none"
//     },
//     idle: {
//       pointerEvents: "auto"
//     }
//   });

// export interface RenderProps {
//   isScrolling: boolean;
//   isAtTheBottom: boolean;
// }

// interface Props {
//   hasResults: boolean;
//   bookmarks: UserBookmark[];
// }

// type ;

interface Props extends ListProps {
  isLoading: boolean;
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
  fetchMore,
  hasResults,
  bookmarks
}: Props) {
  const isAtTheBottom = useWindoBottom();

  // console.log(`isScrolling ${isScrolling}`);
  // const isAtTheBottom = hasReachedTheBottom();
  const offset = bookmarks.length;
  // console.log(`Bottom ${isAtTheBottom}`);

  useEffect(() => {
    // console.log("fetch more");
    // console.log(`Bottom ${isAtTheBottom}`);
    // console.log(`Loading ${isLoading}`);
    // console.log(`Offset ${offset}`);

    if (isAtTheBottom && !isLoading && offset > 0) {
      fetchMore({
        variables: {
          ...variables,
          offset: offset
        },
        updateQuery: (prev, { fetchMoreResult }) => {
          console.log("update query");
          console.log(prev);
          console.log(fetchMoreResult);

          if (!fetchMoreResult) {
            return prev;
          }

          return {
            GetLatestBookmarks: {
              ...fetchMoreResult.GetLatestBookmarks,
              results: [
                ...prev.GetLatestBookmarks.results,
                ...fetchMoreResult.GetLatestBookmarks.results
              ]
            }
          };
        }
      });
    }
  }, [offset, fetchMore, isAtTheBottom, isLoading]);

  return (
    <>
      {isLoading && <Loader />}
      {!isLoading && hasResults && (
        <List hasResults={hasResults} bookmarks={bookmarks} />
      )}
    </>
  );
}
