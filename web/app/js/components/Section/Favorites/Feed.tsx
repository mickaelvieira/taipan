import React, { useEffect, useRef } from "react";
import { ApolloQueryResult } from "apollo-boost";
import Item from "./Item";
import Loader from "../../ui/Loader";
import FavoritesQuery, {
  query,
  variables,
  Data
} from "../../apollo/Query/Favorites";
import { Bookmark } from "../../../types/bookmark";
import FeedContainer from "../../ui/Feed/Container";
import useWindowBottom from "../../../hooks/window-bottom";

type FetchMore = () => Promise<ApolloQueryResult<Data>>;

function hasReceivedData(data: Data | undefined): [boolean, Bookmark[]] {
  let hasResults = false;
  let results: Bookmark[] = [];

  if (data && "GetFavorites" in data && "results" in data.GetFavorites) {
    results = data.GetFavorites.results;
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
    <FavoritesQuery
      query={query}
      variables={variables}
      fetchPolicy="cache-and-network"
    >
      {({ data, loading, fetchMore }) => {
        const [hasResults, bookmarks] = hasReceivedData(data);
        console.log("favorites");
        console.log(data);
        console.log(bookmarks);
        console.log(hasResults);
        console.log(loading);

        if (hasResults) {
          loadMore.current =
            data && data.GetFavorites.results.length === data.GetFavorites.total
              ? undefined
              : () =>
                  fetchMore({
                    variables: {
                      offset: data ? data.GetFavorites.results.length : 0
                    },
                    updateQuery: (prev, { fetchMoreResult }) => {
                      if (!fetchMoreResult) {
                        return prev;
                      }
                      return {
                        GetFavorites: {
                          ...fetchMoreResult.GetFavorites,
                          results: [
                            ...prev.GetFavorites.results,
                            ...fetchMoreResult.GetFavorites.results
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
    </FavoritesQuery>
  );
}
