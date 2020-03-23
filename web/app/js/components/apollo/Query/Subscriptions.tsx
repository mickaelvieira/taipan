import {
  ApolloQueryResult,
  FetchMoreQueryOptions,
  FetchMoreOptions,
} from "apollo-client";
import query from "../graphql/query/subscriptions/subscriptions.graphql";
import { OffsetPagination } from "../../../types";
import { SearchParams } from "../../../types/subscription";
import { SubscriptionResults } from "../../../types/subscription";

export interface Data {
  subscriptions: {
    subscriptions: SubscriptionResults;
  };
}

export interface Variables {
  pagination: OffsetPagination;
  search?: SearchParams;
}

export type FetchMore = <K extends keyof Variables>(
  fetchMoreOptions: FetchMoreQueryOptions<Variables, K> &
    FetchMoreOptions<Data, Variables>
) => Promise<ApolloQueryResult<Data>>;

export type LoadMore = () => Promise<ApolloQueryResult<Data>>;

const variables = {
  pagination: {
    limit: 50,
  },
};

export function getFetchMore(
  fetchMore: FetchMore,
  data: Data,
  variables: Variables
): LoadMore | undefined {
  const {
    subscriptions: {
      subscriptions: { results, total },
    },
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
              subscriptions: {
                ...prev.subscriptions,
                subscriptions: {
                  ...prev.subscriptions.subscriptions,
                  limit: next.subscriptions.subscriptions.limit,
                  offset: next.subscriptions.subscriptions.offset,
                  results: [
                    ...prev.subscriptions.subscriptions.results,
                    ...next.subscriptions.subscriptions.results,
                  ],
                },
              },
            };
          },
        });
}

export { query, variables };
