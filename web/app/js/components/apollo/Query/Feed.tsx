import { Query } from "react-apollo";
import {
  ApolloQueryResult,
  FetchMoreQueryOptions,
  FetchMoreOptions
} from "apollo-boost";
import { FeedVariables, FeedData } from "../../../types/feed";
import { getDataKey } from "../helpers/data";
import queryNews from "../../../services/apollo/query/news.graphql";
import queryReadingList from "../../../services/apollo/query/reading-list.graphql";
import queryFavorites from "../../../services/apollo/query/favorites.graphql";

export type FetchMore = <K extends keyof FeedVariables>(
  fetchMoreOptions: FetchMoreQueryOptions<FeedVariables, K> &
    FetchMoreOptions<FeedData, FeedVariables>
) => Promise<ApolloQueryResult<FeedData>>;

export type LoadMore = () => Promise<ApolloQueryResult<FeedData>>;

const variables = {
  pagination: {
    limit: 10
  }
};

function getFetchMore(
  fetchMore: FetchMore,
  data: FeedData | undefined
): LoadMore | undefined {
  if (!data) {
    return undefined;
  }

  const key = getDataKey(data);

  if (!key) {
    return undefined;
  }

  return data[key].results.length === data[key].total
    ? undefined
    : () =>
        fetchMore({
          variables: {
            pagination: { from: data ? data[key].last : "" }
          },
          updateQuery: (prev, { fetchMoreResult: next }) => {
            if (!next) {
              return prev;
            }

            return {
              [key]: {
                ...prev[key],
                last: next[key].last,
                limit: next[key].limit,
                results: [...prev[key].results, ...next[key].results]
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

class FeedQuery extends Query<FeedData, FeedVariables> {
  static defaultProps = {
    fetchPolicy: "cache-first",
    variables
  };
}

export default FeedQuery;
