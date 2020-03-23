import {
  ApolloQueryResult,
  FetchMoreQueryOptions,
  FetchMoreOptions,
} from "apollo-client";
import PropTypes from "prop-types";
import queryDocuments from "../graphql/query/documents/search.graphql";
import queryBookmarks from "../graphql/query/bookmarks/search.graphql";
import { getDataKey } from "../helpers/search";
import { SearchQueryData, SearchQueryVariables } from "../../../types/search";

export type FetchMore = <K extends keyof SearchQueryVariables>(
  fetchMoreOptions: FetchMoreQueryOptions<SearchQueryVariables, K> &
    FetchMoreOptions<SearchQueryData, SearchQueryVariables>
) => Promise<ApolloQueryResult<SearchQueryData>>;

export type LoadMore = () => Promise<ApolloQueryResult<SearchQueryData>>;

export function getFetchMore(
  fetchMore: FetchMore,
  query: PropTypes.Validator<object>,
  data: SearchQueryData | undefined,
  variables: SearchQueryVariables
): LoadMore | undefined {
  if (!data) {
    return undefined;
  }

  const key = getDataKey(data);

  if (!key) {
    return undefined;
  }

  return data[key].search.results.length === data[key].search.total
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
              [key]: {
                ...prev[key],
                search: {
                  ...prev[key].search,
                  limit: next[key].search.limit,
                  offset: next[key].search.offset,
                  results: [
                    ...prev[key].search.results,
                    ...next[key].search.results,
                  ],
                },
              },
            };
          },
        });
}

const variables = {
  pagination: {
    limit: 50,
  },
};

export { queryDocuments, queryBookmarks, variables };
