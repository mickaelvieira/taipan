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

export interface CursorPagination {
  from?: string;
  to?: string;
  limit?: number;
}

export interface FeedResult {
  total: number;
  first: string;
  last: string;
  limit: number;
  results: FeedItem[];
}

export interface Variables {
  pagination: CursorPagination;
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
  pagination: {
    limit: 10
  }
};

function getDataKey(data: Data): string | null {
  const keys = Object.keys(data);
  return keys.length > 0 ? keys[0] : null;
}

function getBoundaries(results: FeedItem[]): [string, string] {
  let first = "";
  let last = "";
  if (results.length > 0) {
    first = results[0].id;
    last = results[results.length - 1].id;
  }
  return [first, last];
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
  const total = result.total - 1;
  const results = cloned.results.filter(i => i.id !== item.id);
  const [first, last] = getBoundaries(results);

  return {
    ...cloned,
    first,
    last,
    total,
    results
  };
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
