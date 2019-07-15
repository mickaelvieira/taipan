import React, { useEffect, useRef } from "react";
import { withApollo, WithApolloClient } from "react-apollo";
import PropTypes from "prop-types";
import Loader from "../Loader";
import FeedQuery, { LoadMore, getFetchMore } from "../../apollo/Query/Feed";
import FeedSubscription from "../../apollo/Subscription/Feed";
import { FeedItem } from "../../../types/feed";
import {
  FeedUpdater,
  hasReceivedData,
  makeFeedUpdater
} from "../../apollo/helpers/feed";
import FeedContainer from "./Container";
import useWindowBottom from "../../../hooks/window-bottom";

export interface ListProps {
  results: FeedItem[];
  updater: FeedUpdater;
  firstId: string;
  lastId: string;
}

interface Props {
  List: React.FunctionComponent<ListProps>;
  query: PropTypes.Validator<object>;
  subscription: PropTypes.Validator<object>;
}

export default withApollo(function Feed({
  client,
  query,
  subscription,
  List
}: WithApolloClient<Props>): JSX.Element {
  const isAtTheBottom = useWindowBottom();
  const loadMore = useRef<LoadMore | undefined>();
  const updater = makeFeedUpdater(client, query);

  useEffect(() => {
    if (isAtTheBottom && loadMore.current) {
      loadMore.current();
    }
  }, [isAtTheBottom, loadMore]);

  return (
    <>
      <FeedSubscription updater={updater} subscription={subscription} />
      <FeedQuery query={query}>
        {({ data, loading, error, fetchMore }) => {
          const [hasResults, result] = hasReceivedData(data);
          const { results = [], first = "", last = "" } = result;

          if (hasResults) {
            loadMore.current = getFetchMore(fetchMore, data);
          }

          return (
            <>
              {loading && !hasResults && <Loader />}
              {error && !hasResults && <span>{error.message}</span>}
              <FeedContainer>
                <List
                  updater={updater}
                  results={results}
                  firstId={first}
                  lastId={last}
                />
              </FeedContainer>
              {loading && hasResults && <Loader />}
              {error && hasResults && <span>{error.message}</span>}
            </>
          );
        }}
      </FeedQuery>
    </>
  );
});
