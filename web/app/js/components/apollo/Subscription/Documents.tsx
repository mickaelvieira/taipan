import React, { useContext } from "react";
import {
  Subscription as SubscriptionBase,
  OnSubscriptionDataOptions
} from "react-apollo";
import subscription from "../graphql/subscription/documents.graphql";
import { FeedEventData, FeedEvent } from "../../../types/feed";
import { isEmitter } from "../helpers/events";
import { hasReceivedEvent } from "../helpers/feed";
import FeedsUpdater from "../helpers/feeds-updater";
import { ClientContext } from "../../context";
import { Document } from "../../../types/document";

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
      case "unbookmark":
        updater.unbookmark(item as Document);
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
