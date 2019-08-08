import React, { useContext, PropsWithChildren } from "react";
import { useApolloClient, useSubscription } from "@apollo/react-hooks";
import { FeedsCacheContext, FeedsContext, ClientContext } from "../context";
import { FeedEvent, FeedEventData } from "../../types/feed";
import { Document } from "../../types/document";
import { Bookmark } from "../../types/bookmark";
import { isEmitter } from "../apollo/helpers/events";
import { hasReceivedEvent } from "../apollo/helpers/feed";
import FeedsUpdater from "../apollo/helpers/feeds-updater";
import FeedsMutator from "../apollo/helpers/feeds-mutator";
import { documentSubscription } from "../apollo/Subscription/Documents";
import { bookmarkSubscription } from "../apollo/Subscription/Bookmarks";

export default function AppFeeds({
  children
}: PropsWithChildren<{}>): JSX.Element {
  const clientId = useContext(ClientContext);
  const client = useApolloClient();
  const updater = new FeedsUpdater(client);
  const mutator = new FeedsMutator(client, updater);

  useSubscription<FeedEventData>(documentSubscription, {
    onSubscriptionData: ({ subscriptionData }) => {
      const [isReceived, event] = hasReceivedEvent(subscriptionData.data);
      if (isReceived && !isEmitter(event, clientId)) {
        console.log("feed event received and processed");
        console.log(event);
        console.log(clientId);
        const { item, action } = event as FeedEvent;
        switch (action) {
          case "unbookmark":
            updater.unbookmark(item as Document);
            break;
        }
      }
    }
  });
  useSubscription<FeedEventData>(bookmarkSubscription, {
    onSubscriptionData: ({ subscriptionData }) => {
      const [isReceived, event] = hasReceivedEvent(subscriptionData.data);
      if (isReceived && !isEmitter(event, clientId)) {
        console.log("feed event received and processed");
        console.log(event);
        console.log(clientId);
        const { item, action } = event as FeedEvent;
        switch (action) {
          case "bookmark":
            updater.bookmark(item as Bookmark);
            break;
          case "favorite":
            updater.favorite(item as Bookmark);
            break;
          case "unfavorite":
            updater.unfavorite(item as Bookmark);
            break;
        }
      }
    }
  });

  return (
    <FeedsCacheContext.Provider value={updater}>
      <FeedsContext.Provider value={mutator}>{children}</FeedsContext.Provider>
    </FeedsCacheContext.Provider>
  );
}
