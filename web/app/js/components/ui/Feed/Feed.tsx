import React, { useEffect, useRef } from "react";
import PropTypes from "prop-types";
import Loader from "../Loader";
import FeedQuery, { LoadMore, getFetchMore } from "../../apollo/Query/Feed";
import FeedSubscription from "../../apollo/Subscription/Feed";
import { FeedItem } from "../../../types/feed";
import { hasReceivedData } from "../../apollo/helpers/feed";
import FeedContainer from "./Container";
import useWindowBottom from "../../../hooks/window-bottom";

export interface ListProps {
  results: FeedItem[];
  firstId: string;
  lastId: string;
  query: PropTypes.Validator<object>;
}

interface Props {
  List: React.ComponentType<ListProps>;
  query: PropTypes.Validator<object>;
  subscription: PropTypes.Validator<object>;
}

export default function Feed({
  query,
  subscription,
  List
}: Props): JSX.Element {
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
          const [hasResults, result] = hasReceivedData(data);
          const { results = [], first = "", last = "" } = result;

          if (hasResults) {
            loadMore.current = getFetchMore(fetchMore, data);
          }

          return (
            <>
              {loading && !hasResults && <Loader />}
              <FeedContainer>
                <List
                  results={results}
                  firstId={first}
                  lastId={last}
                  query={query}
                />
              </FeedContainer>
              {loading && hasResults && <Loader />}
            </>
          );
        }}
      </FeedQuery>
    </>
  );
}
