import React, { useEffect, useRef } from "react";
import PropTypes from "prop-types";
import Loader from "../Loader";
import FeedQuery, { LoadMore, getFetchMore } from "../../apollo/Query/Feed";
import FeedSubscription from "../../apollo/Subscription/Feed";
import { FeedItem } from "../../../types/feed";
import { hasReceivedData } from "../../apollo/helpers/data";
import FeedContainer from "./Container";
import useWindowBottom from "../../../hooks/window-bottom";

export interface ListProps {
  results: FeedItem[];
  query: PropTypes.Validator<object>;
}

interface Props {
  List: React.ComponentType<ListProps>;
  query: PropTypes.Validator<object>;
  subscription: PropTypes.Validator<object>;
}

export default function Feed({ query, subscription, List }: Props) {
  const isAtTheBottom = useWindowBottom();
  const loadMore = useRef<LoadMore | undefined>();

  useEffect(() => {
    if (isAtTheBottom && loadMore.current) {
      loadMore.current();
    }
  }, [isAtTheBottom, loadMore]);

  return (
    <>
      <FeedSubscription query={query} subscription={subscription} />
      <FeedQuery query={query}>
        {({ data, loading, fetchMore }) => {
          const [hasResults, results] = hasReceivedData(data);

          if (hasResults) {
            loadMore.current = getFetchMore(fetchMore, data);
          }

          return (
            <>
              {loading && !hasResults && <Loader />}
              <FeedContainer>
                <List results={results} query={query} />
              </FeedContainer>
              {loading && hasResults && <Loader />}
            </>
          );
        }}
      </FeedQuery>
    </>
  );
}
