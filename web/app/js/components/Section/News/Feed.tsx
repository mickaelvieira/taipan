import React, { useEffect, useRef } from "react";
import { ApolloQueryResult } from "apollo-boost";
import Item from "./Item";
import Loader from "../../ui/Loader";
import NewsQuery, { query, variables, Data } from "../../apollo/Query/News";
import { Document } from "../../../types/document";
import FeedContainer from "../../ui/Feed/Container";
import useWindowBottom from "../../../hooks/window-bottom";

type FetchMore = () => Promise<ApolloQueryResult<Data>>;

function hasReceivedData(data: Data | undefined): [boolean, Document[]] {
  let hasResults = false;
  let results: Document[] = [];

  if (data && "News" in data && "results" in data.News) {
    results = data.News.results;
    if (results.length > 0) {
      hasResults = true;
    }
  }

  return [hasResults, results];
}

export default function Feed() {
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
    <NewsQuery query={query} variables={variables}>
      {({ data, loading, fetchMore }) => {
        const [hasResults, documents] = hasReceivedData(data);
        console.log("news");
        console.log(data);
        console.log(document);
        console.log(hasResults);
        console.log(loading);

        if (hasResults) {
          loadMore.current =
            data && data.News.results.length === data.News.total
              ? undefined
              : () =>
                  fetchMore({
                    variables: {
                      offset: data ? data.News.results.length : 0
                    },
                    updateQuery: (prev, { fetchMoreResult }) => {
                      if (!fetchMoreResult) {
                        return prev;
                      }
                      return {
                        News: {
                          ...fetchMoreResult.News,
                          results: [
                            ...prev.News.results,
                            ...fetchMoreResult.News.results
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
              {documents.map((document: Document, index) => (
                <Item document={document} index={index} key={document.id} />
              ))}
            </FeedContainer>
            {loading && hasResults && <Loader />}
          </>
        );
      }}
    </NewsQuery>
  );
}
