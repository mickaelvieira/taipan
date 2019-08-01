import { Query } from "react-apollo";
import {
  ApolloQueryResult,
  FetchMoreQueryOptions,
  FetchMoreOptions
} from "apollo-boost";
import query from "../graphql/query/documents/search.graphql";
import { OffsetPagination } from "../../../types";
import { SearchQueryData } from "../../../types/search";
import { SearchParams } from "../../../types/document";

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
    documents: {
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
              documents: {
                ...prev.documents,
                search: {
                  ...prev.documents.search,
                  limit: next.documents.search.limit,
                  offset: next.documents.search.offset,
                  results: [
                    ...prev.documents.search.results,
                    ...next.documents.search.results
                  ]
                }
              }
            };
          }
        });
}

const variables = {
  pagination: {
    limit: 50
  }
};

export { query, variables };

class DocumentsQuery extends Query<Data, Variables> {
  static defaultProps = {
    query,
    variables
  };
}

export default DocumentsQuery;
