import React, { useContext } from "react";
import {
  Subscription as SubscriptionBase,
  OnSubscriptionDataOptions
} from "react-apollo";
import PropTypes from "prop-types";
import subscriptionNews from "../graphql/subscription/news.graphql";
import subscriptionFavorites from "../graphql/subscription/favorites.graphql";
import subscriptionReadingList from "../graphql/subscription/reading-list.graphql";
import { getDataKey } from "../Query/Feed";
import { FeedEventData, FeedEvent, FeedQueryData } from "../../../types/feed";
import { hasReceivedEvent, feedResultsAction } from "../helpers/feed";
import { ClientContext } from "../../context";

interface Data extends OnSubscriptionDataOptions<FeedEventData> {
  query: PropTypes.Validator<object>;
  clientId: string;
}

function isEmitter(event: FeedEvent | null, clientId: string): boolean {
  if (!event) {
    return false;
  }
  return event.emitter === clientId;
}

function onReceivedData({
  client,
  subscriptionData,
  query,
  clientId
}: Data): void {
  const [isReceived, event] = hasReceivedEvent(subscriptionData.data);
  console.log(event);
  console.log(clientId);
  if (isReceived && !isEmitter(event, clientId)) {
    const { item, action } = event as FeedEvent;
    const data = client.readQuery({ query }) as FeedQueryData;
    const updateResult = feedResultsAction[action];
    if (data) {
      const key = getDataKey(data);
      if (key) {
        const result = updateResult(data.feeds[key], item);
        client.writeQuery({
          query,
          data: {
            feeds: {
              ...data.feeds,
              [key]: result
            }
          }
        });
      }
    }
  }
}

export { subscriptionNews, subscriptionFavorites, subscriptionReadingList };

class Subscription extends SubscriptionBase<FeedEventData, {}> {}

interface Props {
  query: PropTypes.Validator<object>;
  subscription: PropTypes.Validator<object>;
}

export default function FeedSubscription({
  subscription,
  query
}: Props): JSX.Element {
  const clientId = useContext(ClientContext);
  return (
    <Subscription
      subscription={subscription}
      onSubscriptionData={options =>
        onReceivedData({ ...options, query, clientId })
      }
    />
  );
}
