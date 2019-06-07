import { cloneDeep } from "lodash";
import { Query } from "react-apollo";
import {
  ApolloQueryResult,
  FetchMoreQueryOptions,
  FetchMoreOptions
} from "apollo-boost";
import { Bookmark } from "../../../types/bookmark";
import { Document } from "../../../types/document";
import queryNews from "../../../services/apollo/query/news.graphql";
import queryReadingList from "../../../services/apollo/query/reading-list.graphql";
import queryFavorites from "../../../services/apollo/query/favorites.graphql";

export type FeedItem = Bookmark | Document;

export interface FeedResult {
  total: number;
  offset: number;
  limit: number;
  results: FeedItem[];
}

export interface Variables {
  offset?: number;
  limit?: number;
}

export interface Data {
  [key: string]: FeedResult;
}

export type FetchMore = <K extends keyof Variables>(
  fetchMoreOptions: FetchMoreQueryOptions<Variables, K> &
    FetchMoreOptions<Data, Variables>
) => Promise<ApolloQueryResult<Data>>;

export type LoadMore = () => Promise<ApolloQueryResult<Data>>;

const variables = {
  limit: 10
};

function getDataKey(data: Data): string | null {
  const keys = Object.keys(data);
  return keys.length > 0 ? keys[0] : null;
}

const removeItemFromFeedResults = (
  result: FeedResult,
  item: FeedItem
): FeedResult => {
  if (!item) {
    return result;
  }

  const index = result.results.findIndex(i => i.id === item.id);
  if (index < 0) {
    return result;
  }

  const cloned = cloneDeep(result);
  cloned.total = result.total - 1;
  cloned.results = cloned.results.filter(i => i.id !== item.id);

  return cloned;
};

function hasReceivedData(data: Data | undefined): [boolean, FeedItem[]] {
  let hasResults = false;
  let results: FeedItem[] = [];

  if (data) {
    const key = getDataKey(data);
    if (key && "results" in data[key]) {
      results = data[key].results;
      if (results.length > 0) {
        hasResults = true;
      }
    }
  }

  return [hasResults, results];
}

function getFetchMore(
  fetchMore: FetchMore,
  data: Data | undefined
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
            offset: data ? data[key].results.length : 0
          },
          updateQuery: (prev, { fetchMoreResult }) => {
            if (!fetchMoreResult) {
              return prev;
            }
            return {
              [key]: {
                ...fetchMoreResult[key],
                results: [...prev[key].results, ...fetchMoreResult[key].results]
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
  removeItemFromFeedResults,
  hasReceivedData,
  getFetchMore,
  getDataKey
};

class FeedQuery extends Query<Data, Variables> {
  static defaultProps = {
    fetchPolicy: "cache-first",
    variables
  };
}

export default FeedQuery;
