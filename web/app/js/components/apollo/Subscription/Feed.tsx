import React, { useContext } from "react";
import {
  Subscription as SubscriptionBase,
  OnSubscriptionDataOptions
} from "react-apollo";
import PropTypes from "prop-types";
import subscriptionNews from "../graphql/subscription/news.graphql";
import subscriptionFavorites from "../graphql/subscription/favorites.graphql";
import subscriptionReadingList from "../graphql/subscription/reading-list.graphql";
import { FeedEventData, FeedEvent } from "../../../types/feed";
import { hasReceivedEvent, FeedUpdater } from "../helpers/feed";
import { ClientContext } from "../../context";

interface Data extends OnSubscriptionDataOptions<FeedEventData> {
  updater: FeedUpdater;
  clientId: string;
}

function isEmitter(event: FeedEvent | null, clientId: string): boolean {
  if (!event) {
    return false;
  }
  return event.emitter === clientId;
}

function onReceivedData({ subscriptionData, updater, clientId }: Data): void {
  const [isReceived, event] = hasReceivedEvent(subscriptionData.data);
  console.log(event);
  console.log(clientId);
  if (isReceived && !isEmitter(event, clientId)) {
    const { item, action } = event as FeedEvent;
    updater(item, action);
  }
}

export { subscriptionNews, subscriptionFavorites, subscriptionReadingList };

class Subscription extends SubscriptionBase<FeedEventData, {}> {}

interface Props {
  updater: FeedUpdater;
  subscription: PropTypes.Validator<object>;
}

export default function FeedSubscription({
  subscription,
  updater
}: Props): JSX.Element {
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
