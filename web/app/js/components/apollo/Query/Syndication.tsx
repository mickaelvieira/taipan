import {
  ApolloQueryResult,
  FetchMoreQueryOptions,
  FetchMoreOptions
} from "apollo-client";
import query from "../graphql/query/syndication/sources.graphql";
import { OffsetPagination } from "../../../types";
import { SyndicationResults, SearchParams } from "../../../types/syndication";

export interface Data {
  syndication: {
    sources: SyndicationResults;
  };
}

export interface Variables {
  pagination: OffsetPagination;
  search: SearchParams;
}

const variables = {
  pagination: {
    limit: 50
  }
};

export type FetchMore = <K extends keyof Variables>(
  fetchMoreOptions: FetchMoreQueryOptions<Variables, K> &
    FetchMoreOptions<Data, Variables>
) => Promise<ApolloQueryResult<Data>>;

export type LoadMore = () => Promise<ApolloQueryResult<Data>>;

export function getFetchMore(
  fetchMore: FetchMore,
  data: Data,
  variables: Variables
): LoadMore | undefined {
  const {
    syndication: {
      sources: { results, total }
    }
  } = data;

  return results.length === total
    ? undefined
    : () =>
        fetchMore({
          query,
          variables,
          updateQuery: (prev: Data, { fetchMoreResult: next }) => {
            if (!next) {
              return prev;
            }
            return {
              syndication: {
                ...prev.syndication,
                sources: {
                  ...prev.syndication.sources,
                  limit: next.syndication.sources.limit,
                  offset: next.syndication.sources.offset,
                  results: [
                    ...prev.syndication.sources.results,
                    ...next.syndication.sources.results
                  ]
                }
              }
            };
          }
        });
}

export { query, variables };
