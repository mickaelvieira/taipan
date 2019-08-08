import {
  ApolloQueryResult,
  FetchMoreQueryOptions,
  FetchMoreOptions
} from "apollo-client";
import query from "../graphql/query/bookmarks/search.graphql";
import { OffsetPagination } from "../../../types";
import { SearchQueryData } from "../../../types/search";
import { SearchParams } from "../../../types/bookmark";

export interface Variables {
  pagination: OffsetPagination;
  search?: SearchParams;
}

export type FetchMore = <K extends keyof Variables>(
  fetchMoreOptions: FetchMoreQueryOptions<Variables, K> &
    FetchMoreOptions<SearchQueryData, Variables>
) => Promise<ApolloQueryResult<SearchQueryData>>;

export type LoadMore = () => Promise<ApolloQueryResult<SearchQueryData>>;

export function getFetchMore(
  fetchMore: FetchMore,
  data: SearchQueryData | undefined,
  variables: Variables
): LoadMore | undefined {
  if (!data) {
    return undefined;
  }

  const {
    bookmarks: {
      search: { results, total }
    }
  } = data;

  return results.length === total
    ? undefined
    : () =>
        fetchMore({
          query,
          variables,
          updateQuery: (prev: SearchQueryData, { fetchMoreResult: next }) => {
            if (!next) {
              return prev;
            }
            return {
              bookmarks: {
                ...prev.bookmarks,
                search: {
                  ...prev.bookmarks.search,
                  limit: next.bookmarks.search.limit,
                  offset: next.bookmarks.search.offset,
                  results: [
                    ...prev.bookmarks.search.results,
                    ...next.bookmarks.search.results
                  ]
                }
              }
            };
          }
        });
}

const variables = {
  pagination: {
    limit: 20
  }
};

export { query, variables };
