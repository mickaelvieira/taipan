import React, { useEffect, ReactNode } from "react";
import { connect } from "react-redux";
import { Dispatch } from "redux";
import { pushFeedItems } from "store/actions/feed";
import { ReduxAction } from "store/actions/types";
import { RootState } from "store/reducer/default";
import { Bookmark, BookmarksLinks } from "types/bookmark";
import { CollectionResponse } from "collection/types";
import useWorkers from "../hooks/workers";
import useConnectionStatus from "../hooks/connection-status";

interface Props {
  isNearTheEnd: boolean;
  links: BookmarksLinks;
  pushFeedItems: (item: Bookmark[]) => void;
  children: ReactNode;
}

const InfiniteFeed = (props: Props) => {
  const { links, isNearTheEnd, pushFeedItems, children } = props;
  const isOnline = useConnectionStatus();
  const workers = useWorkers(pushFeedItems);

  useEffect(() => {
    if (isOnline && isNearTheEnd && links.next) {
      workers.postMessage(links.next);
    }
  });

  return <>{children}</>;
};

const mapStateToProps = (state: RootState) => ({
  links: state.feed.links
});

const mapDispatchToProps = (dispatch: Dispatch<ReduxAction>) => ({
  pushFeedItems: (json: CollectionResponse) => dispatch(pushFeedItems(json))
});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(InfiniteFeed);
