import React, { PropsWithChildren } from "react";
import { FeedsCacheContext, FeedsContext } from "../context";
import { ApolloClient } from "apollo-client";
import FeedsUpdater from "../apollo/helpers/feeds-updater";
import FeedsMutator from "../apollo/helpers/feeds-mutator";
import DocumentsSubscription from "../apollo/Subscription/Documents";
import BookmarksSubscription from "../apollo/Subscription/Bookmarks";

interface Props {
  client: ApolloClient<object>;
}

export default function AppFeeds({
  children,
  client
}: PropsWithChildren<Props>): JSX.Element {
  const updater = new FeedsUpdater(client);
  const mutator = new FeedsMutator(client, updater);

  return (
    <FeedsCacheContext.Provider value={updater}>
      <FeedsContext.Provider value={mutator}>
        <DocumentsSubscription updater={updater} />
        <BookmarksSubscription updater={updater} />
        {children}
      </FeedsContext.Provider>
    </FeedsCacheContext.Provider>
  );
}
