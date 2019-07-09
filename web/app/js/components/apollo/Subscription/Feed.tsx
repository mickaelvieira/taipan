import React from "react";
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

interface Data extends OnSubscriptionDataOptions<FeedEventData> {
  query: PropTypes.Validator<object>;
}

function onReceivedData({ client, subscriptionData, query }: Data): void {
  const [isReceived, event] = hasReceivedEvent(subscriptionData.data);
  console.log(event);
  if (isReceived) {
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

class Subscription extends SubscriptionBase<FeedEventData, {}> { }

interface Props {
  query: PropTypes.Validator<object>;
  subscription: PropTypes.Validator<object>;
}

export default function FeedSubscription({
  subscription,
  query
}: Props): JSX.Element {
  return (
    <Subscription
      subscription={subscription}
      onSubscriptionData={options => onReceivedData({ ...options, query })}
    />
  );
}
