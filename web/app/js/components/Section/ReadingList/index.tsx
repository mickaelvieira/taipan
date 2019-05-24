import React, { useEffect, useRef } from "react";
import { ApolloQueryResult } from "apollo-boost";
import Item from "./Item";
import Loader from "../../ui/Loader";
import ReadingListQuery, {
  query,
  variables,
  Data
} from "../../apollo/Query/ReadingList";
import { Bookmark } from "../../../types/bookmark";
import FeedContainer from "../../ui/Feed/Container";
import useWindowBottom from "../../../hooks/window-bottom";

type FetchMore = () => Promise<ApolloQueryResult<Data>>;

function hasReceivedData(data: Data | undefined): [boolean, Bookmark[]] {
  let hasResults = false;
  let results: Bookmark[] = [];

  if (data && "GetReadingList" in data && "results" in data.GetReadingList) {
    results = data.GetReadingList.results;
    if (results.length > 0) {
      hasResults = true;
    }
  }

  return [hasResults, results];
}

export default function ReadingList() {
  const isAtTheBottom = useWindowBottom();
  const loadMore = useRef<FetchMore | undefined>();

  console.log("func");
  console.log(isAtTheBottom);

  useEffect(() => {
    console.log("effect");
    console.log(isAtTheBottom);

    if (isAtTheBottom && loadMore.current) {
      console.log("======================= FETCH MORE =======================");
      loadMore.current();
    }
  }, [isAtTheBottom, loadMore]);

  return (
    <ReadingListQuery query={query} variables={variables}>
      {({ data, loading, fetchMore }) => {
        const [hasResults, bookmarks] = hasReceivedData(data);
        console.log("reading list");
        console.log(data);
        console.log(bookmarks);
        console.log(hasResults);
        console.log(loading);

        if (hasResults) {
          loadMore.current =
            data &&
            data.GetReadingList.results.length === data.GetReadingList.total
              ? undefined
              : () =>
                  fetchMore({
                    variables: {
                      offset: data ? data.GetReadingList.results.length : 0
                    },
                    updateQuery: (prev, { fetchMoreResult }) => {
                      if (!fetchMoreResult) {
                        return prev;
                      }
                      return {
                        GetReadingList: {
                          ...fetchMoreResult.GetReadingList,
                          results: [
                            ...prev.GetReadingList.results,
                            ...fetchMoreResult.GetReadingList.results
                          ]
                        }
                      };
                    }
                  });
        }

        return (
          <>
            {loading && !hasResults && <Loader />}
            <FeedContainer>
              {bookmarks.map((bookmark: Bookmark, index) => (
                <Item bookmark={bookmark} index={index} key={bookmark.id} />
              ))}
            </FeedContainer>
            {loading && hasResults && <Loader />}
          </>
        );
      }}
    </ReadingListQuery>
  );
}
