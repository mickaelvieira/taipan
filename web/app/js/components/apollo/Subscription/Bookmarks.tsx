import React, { useContext } from "react";
import {
  Subscription as SubscriptionBase,
  OnSubscriptionDataOptions
} from "react-apollo";
import subscription from "../graphql/subscription/bookmarks.graphql";
import { FeedEventData, FeedEvent } from "../../../types/feed";
import { isEmitter } from "../helpers/events";
import { hasReceivedEvent } from "../helpers/feed";
import FeedsUpdater from "../helpers/feeds-updater";
import { ClientContext } from "../../context";
import { Bookmark } from "../../../types/bookmark";

interface Data extends OnSubscriptionDataOptions<FeedEventData> {
  updater: FeedsUpdater;
  clientId: string;
}

function onReceivedData({ subscriptionData, updater, clientId }: Data): void {
  const [isReceived, event] = hasReceivedEvent(subscriptionData.data);
  console.log(event);
  console.log(clientId);
  if (isReceived && !isEmitter(event, clientId)) {
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

export { subscription };

class Subscription extends SubscriptionBase<FeedEventData, {}> {}

interface Props {
  updater: FeedsUpdater;
}

export default function FeedSubscription({ updater }: Props): JSX.Element {
  const clientId = useContext(ClientContext);
  return (
    <Subscription
      subscription={subscription}
      onSubscriptionData={options =>
        onReceivedData({ ...options, updater, clientId })
      }
    />
  );
}
