import React, { useEffect, useRef } from "react";
import { withApollo, WithApolloClient } from "react-apollo";
import PropTypes from "prop-types";
import Loader from "../Loader";
import FeedQuery, { LoadMore, getFetchMore } from "../../apollo/Query/Feed";
import { FeedItem } from "../../../types/feed";
import { hasReceivedData } from "../../apollo/helpers/feed";
import FeedContainer from "./Container";
import useWindowBottom from "../../../hooks/window-bottom";

export interface ListProps {
  results: FeedItem[];
  firstId: string;
  lastId: string;
}

interface Props {
  List: React.FunctionComponent<ListProps>;
  query: PropTypes.Validator<object>;
}

export default withApollo(function Feed({
  query,
  List
}: WithApolloClient<Props>): JSX.Element {
  const isAtTheBottom = useWindowBottom();
  const loadMore = useRef<LoadMore | undefined>();

  useEffect(() => {
    if (isAtTheBottom && loadMore.current) {
      loadMore.current();
    }
  }, [isAtTheBottom, loadMore]);

  return (
    <FeedQuery query={query}>
      {({ data, loading, error, fetchMore, networkStatus }) => {
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
              <FeedContainer>
                <List results={results} firstId={first} lastId={last} />
              </FeedContainer>
            )}
            {isFetchingMore && hasResults && <Loader />}
            {error && hasResults && <span>{error.message}</span>}
          </>
        );
      }}
    </FeedQuery>
  );
});
