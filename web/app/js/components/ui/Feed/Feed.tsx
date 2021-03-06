import React, { useEffect, useRef } from "react";
import { useQuery } from "@apollo/react-hooks";
import PropTypes from "prop-types";
import Loader from "../Loader";
import { LoadMore, getFetchMore, variables } from "../../apollo/Query/Feed";
import {
  FeedItem,
  FeedVariables,
  FeedQueryData,
  FeedName,
} from "../../../types/feed";
import { hasReceivedData } from "../../apollo/helpers/feed";
import FeedContainer from "./Container";
import useWindowBottom from "../../../hooks/useWindowBottom";

export interface ListProps {
  results: FeedItem[];
  firstId?: string;
  lastId?: string;
}

interface Props {
  name: FeedName;
  List: React.FunctionComponent<ListProps>;
  query: PropTypes.Validator<object>;
}

export default function Feed({ query, List, name }: Props): JSX.Element {
  const isAtTheBottom = useWindowBottom(1000);
  const loadMore = useRef<LoadMore | undefined>();
  const { loading, error, data, fetchMore, networkStatus } = useQuery<
    FeedQueryData,
    FeedVariables
  >(query, {
    variables,
    // notifyOnNetworkStatusChange: true
  });

  useEffect(() => {
    if (isAtTheBottom && loadMore.current) {
      loadMore.current();
    }
  }, [isAtTheBottom, loadMore]);

  const [hasResults, result] = hasReceivedData(data);
  const { results = [], first = "", last = "" } = result;
  const isFetchingFirst = loading && networkStatus === 1;
  const isFetchingMore = loading && networkStatus === 3;

  if (hasResults) {
    loadMore.current = getFetchMore(fetchMore, data);
  }

  return (
    <>
      {isFetchingFirst && !hasResults && <Loader />}
      {error && !hasResults && <span>{error.message}</span>}
      {!isFetchingFirst && !error && (
        <FeedContainer
          name={name}
          List={List}
          results={results}
          firstId={first}
          lastId={last}
        />
      )}
      {isFetchingMore && hasResults && <Loader />}
      {error && hasResults && <span>{error.message}</span>}
    </>
  );
}
