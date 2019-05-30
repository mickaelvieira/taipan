import React, { useEffect, useRef } from "react";
import PropTypes from "prop-types";
import Loader from "../Loader";
import FeedQuery, {
  hasReceivedData,
  getFetchMore,
  variables,
  LoadMore,
  DataKey,
  DataType
} from "../../apollo/Query/Feed";
import FeedContainer from "./Container";
import useWindowBottom from "../../../hooks/window-bottom";

interface ListProps {
  results: DataType;
}

interface Props {
  dataKey: DataKey;
  List: React.ComponentType<ListProps>;
  query: PropTypes.Validator<object>;
}

export default function Feed({ query, dataKey, List }: Props) {
  const isAtTheBottom = useWindowBottom();
  const loadMore = useRef<LoadMore | undefined>();

  useEffect(() => {
    if (isAtTheBottom && loadMore.current) {
      loadMore.current();
    }
  }, [isAtTheBottom, loadMore]);

  return (
    <FeedQuery query={query} variables={variables}>
      {({ data, loading, fetchMore }) => {
        const [hasResults, results] = hasReceivedData(dataKey, data);

        if (hasResults) {
          loadMore.current = getFetchMore(fetchMore, dataKey, data);
        }

        return (
          <>
            {loading && !hasResults && <Loader />}
            <FeedContainer>
              <List results={results} />
            </FeedContainer>
            {loading && hasResults && <Loader />}
          </>
        );
      }}
    </FeedQuery>
  );
}
