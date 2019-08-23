import React, { useRef, useEffect } from "react";
import { useQuery } from "@apollo/react-hooks";
import PropTypes from "prop-types";
import { variables, getFetchMore, LoadMore } from "../../apollo/Query/Search";
import { FeedName } from "../../../types/feed";
import { SearchQueryData, SearchQueryVariables } from "../../../types/search";
import { hasReceivedData } from "../../apollo/helpers/search";
import useWindowBottom from "../../../hooks/useWindowBottom";
import Loader from "../../ui/Loader";
import FeedContainer from "../../ui/Feed/Container";
import { ListProps } from "../../ui/Feed/Feed";
import Pagination from "./Pagination";
import { SearchProps } from ".";

interface Props extends SearchProps {
  name: FeedName;
  List: React.FunctionComponent<ListProps & SearchProps>;
  query: PropTypes.Validator<object>;
}

export default function Feed({
  terms,
  type,
  query,
  List,
  name
}: Props): JSX.Element {
  const isAtTheBottom = useWindowBottom();
  const loadMore = useRef<LoadMore | undefined>();
  const { data, loading, error, networkStatus, fetchMore } = useQuery<
    SearchQueryData,
    SearchQueryVariables
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
    loadMore.current = getFetchMore(fetchMore, query, data, {
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
          <FeedContainer
            name={name}
            List={List}
            results={results}
            terms={terms}
            type={type}
          />
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
