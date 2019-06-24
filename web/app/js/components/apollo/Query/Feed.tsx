import { Query } from "react-apollo";
import {
  ApolloQueryResult,
  FetchMoreQueryOptions,
  FetchMoreOptions
} from "apollo-boost";
import { FeedVariables, FeedQueryData } from "../../../types/feed";
import { getDataKey } from "../helpers/data";
import queryNews from "../../../services/apollo/query/feeds/news.graphql";
import queryReadingList from "../../../services/apollo/query/feeds/reading-list.graphql";
import queryFavorites from "../../../services/apollo/query/feeds/favorites.graphql";

export type FetchMore = <K extends keyof FeedVariables>(
  fetchMoreOptions: FetchMoreQueryOptions<FeedVariables, K> &
    FetchMoreOptions<FeedQueryData, FeedVariables>
) => Promise<ApolloQueryResult<FeedQueryData>>;

export type LoadMore = () => Promise<ApolloQueryResult<FeedQueryData>>;

const variables = {
  pagination: {
    limit: 10
  }
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
            pagination: { from: data ? data.feeds[key].last : "" }
          },
          updateQuery: (prev, { fetchMoreResult: next }) => {
            if (!next) {
              return prev;
            }

            return {
              feeds: {
                [key]: {
                  ...prev.feeds[key],
                  last: next.feeds[key].last,
                  limit: next.feeds[key].limit,
                  results: [
                    ...prev.feeds[key].results,
                    ...next.feeds[key].results
                  ]
                }
              }
            };
          }
        });
}

export {
  queryFavorites,
  queryReadingList,
  queryNews,
  variables,
  getFetchMore,
  getDataKey
};

class FeedQuery extends Query<FeedQueryData, FeedVariables> {
  static defaultProps = {
    fetchPolicy: "cache-first",
    variables
  };
}

export default FeedQuery;
