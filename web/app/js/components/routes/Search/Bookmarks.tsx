import React, { useRef, useEffect } from "react";
import { useQuery } from "@apollo/react-hooks";
import {
  query,
  variables,
  getFetchMore,
  LoadMore,
  Variables
} from "../../apollo/Query/Bookmarks";
import { SearchQueryData } from "../../../types/search";
import { hasReceivedData } from "../../apollo/helpers/search";
import useWindowBottom from "../../../hooks/useWindowBottom";
import Loader from "../../ui/Loader";
import Pagination from "./Pagination";
import Results from "./Results";
import { SearchType } from "../../../types/search";

interface Props {
  type: SearchType;
  terms: string[];
}

export default function SearchBookmarks({ type, terms }: Props): JSX.Element {
  const isAtTheBottom = useWindowBottom();
  const loadMore = useRef<LoadMore | undefined>();
  const { data, loading, error, networkStatus, fetchMore } = useQuery<
    SearchQueryData,
    Variables
  >(query, {
    fetchPolicy: "network-only",
    skip: terms.length === 0,
    variables: { ...variables, search: { terms } }
  });

  useEffect(() => {
    if (isAtTheBottom && loadMore.current) {
      loadMore.current();
    }
  }, [isAtTheBottom, loadMore]);

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
}
