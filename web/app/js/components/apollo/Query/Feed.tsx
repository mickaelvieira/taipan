import {
  ApolloQueryResult,
  FetchMoreQueryOptions,
  FetchMoreOptions,
} from "apollo-client";
import { FeedVariables, FeedQueryData } from "../../../types/feed";
import { getDataKey } from "../helpers/feed";
import queryNews from "../graphql/query/feeds/news.graphql";
import queryReadingList from "../graphql/query/feeds/reading-list.graphql";
import queryFavorites from "../graphql/query/feeds/favorites.graphql";

export type FetchMore = <K extends keyof FeedVariables>(
  fetchMoreOptions: FetchMoreQueryOptions<FeedVariables, K> &
    FetchMoreOptions<FeedQueryData, FeedVariables>
) => Promise<ApolloQueryResult<FeedQueryData>>;

export type LoadMore = () => Promise<ApolloQueryResult<FeedQueryData>>;

const variables = {
  pagination: {
    limit: 10,
  },
};

function getFetchMore(
  fetchMore: FetchMore,
  data: FeedQueryData | undefined
): LoadMore | undefined {
  if (!data) {
    return undefined;
  }

  const key = getDataKey(data);

  if (!key) {
    return undefined;
  }

  return data.feeds[key].results.length === data.feeds[key].total
    ? undefined
    : () =>
        fetchMore({
          variables: {
            pagination: {
              ...variables.pagination,
              from: data ? data.feeds[key].last : "",
            },
          },
          updateQuery: (prev, { fetchMoreResult: next }) => {
            if (!next) {
              return prev;
            }

            return {
              feeds: {
                ...prev.feeds,
                [key]: {
                  ...prev.feeds[key],
                  last: next.feeds[key].last,
                  limit: next.feeds[key].limit,
                  results: [
                    ...prev.feeds[key].results,
                    ...next.feeds[key].results,
                  ],
                },
              },
            };
          },
        });
}

export {
  queryFavorites,
  queryReadingList,
  queryNews,
  variables,
  getFetchMore,
  getDataKey,
};
