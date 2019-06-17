import React from "react";
import { Subscription, OnSubscriptionDataOptions } from "react-apollo";
import PropTypes from "prop-types";
import subscriptionNews from "../../../services/apollo/subscription/news.graphql";
import subscriptionFavorites from "../../../services/apollo/subscription/favorites.graphql";
import subscriptionReadingList from "../../../services/apollo/subscription/reading-list.graphql";
import { getDataKey } from "../Query/Feed";
import { FeedEventData, FeedEvent } from "../../../types/feed";
import { feedResultsAction } from "../helpers/feed";
import { hasReceivedEvent } from "../helpers/data";

const variables = {};

interface Data extends OnSubscriptionDataOptions<FeedEventData> {
  query: PropTypes.Validator<object>;
}

export function onReceivedData({ client, subscriptionData, query }: Data) {
  const [isReceived, event] = hasReceivedEvent(subscriptionData.data);
  console.log(event);
  if (isReceived) {
    const { item, action } = event as FeedEvent;
    const data = client.readQuery({ query });
    const updateResult = feedResultsAction[action];
    if (data) {
      const key = getDataKey(data);
      if (key) {
        const result = updateResult(data[key], item);
        client.writeQuery({
          query,
          data: { [key]: result }
        });
      }
    }
  }
}

export {
  subscriptionNews,
  subscriptionFavorites,
  subscriptionReadingList,
  variables
};

class FeedSubscription extends Subscription<FeedEventData, {}> {}

interface Props {
  query: PropTypes.Validator<object>;
  subscription: PropTypes.Validator<object>;
}

export default function({ subscription, query }: Props) {
  return (
    <FeedSubscription
      subscription={subscription}
      onSubscriptionData={options => onReceivedData({ ...options, query })}
    >
      {() => null}
    </FeedSubscription>
  );
}
