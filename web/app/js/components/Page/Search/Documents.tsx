import React, { useRef, useEffect } from "react";
import DocumentSearchSearch, {
  variables,
  getFetchMore,
  LoadMore
} from "../../apollo/Query/Documents";
import { hasReceivedData } from "../../apollo/helpers/search";
import useWindowBottom from "../../../hooks/window-bottom";
import Loader from "../../ui/Loader";
import Pagination from "./Pagination";
import Results from "./Results";
import { SearchType } from "../../../types/search";

interface Props {
  terms: string[];
}

interface Props {
  type: SearchType;
  terms: string[];
}

export default function SearchDocuments({ terms, type }: Props): JSX.Element {
  const isAtTheBottom = useWindowBottom();
  const loadMore = useRef<LoadMore | undefined>();

  useEffect(() => {
    if (isAtTheBottom && loadMore.current) {
      loadMore.current();
    }
  }, [isAtTheBottom, loadMore]);

  return (
    <DocumentSearchSearch
      fetchPolicy="network-only"
      skip={terms.length === 0}
      variables={{ ...variables, search: { terms } }}
    >
      {({ data, loading, error, networkStatus, fetchMore }) => {
        const [hasResults, result] = hasReceivedData(data);
        const { results = [], total = 0 } = result;
        const isFetchingFirst = loading && networkStatus === 1;
        const isFetchingMore = loading && networkStatus === 3;

        if (hasResults) {
          loadMore.current = getFetchMore(fetchMore, data, {
            ...variables,
            pagination: {
              ...variables.pagination,
              offset: results.length
            },
            search: { terms }
          });
        }

        return (
          <>
            {isFetchingFirst && !hasResults && <Loader />}
            {error && !hasResults && <span>{error.message}</span>}
            {!isFetchingFirst && !error && (
              <>
                <Pagination
                  count={results.length}
                  total={total}
                  terms={terms}
                  type={type}
                />
                <Results results={results} type={type} terms={terms} />
                <Pagination
                  withCount
                  count={results.length}
                  total={total}
                  terms={terms}
                  type={type}
                />
              </>
            )}
            {isFetchingMore && hasResults && <Loader />}
            {error && hasResults && <span>{error.message}</span>}
          </>
        );
      }}
    </DocumentSearchSearch>
  );
}
