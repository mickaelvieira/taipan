import React from "react";
import {
  Subscription as SubscriptionBase,
  OnSubscriptionDataOptions
} from "react-apollo";
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

function onReceivedData({ client, subscriptionData, query }: Data) {
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
  variables,
  subscriptionNews,
  subscriptionFavorites,
  subscriptionReadingList
};

class Subscription extends SubscriptionBase<FeedEventData, {}> {}

interface Props {
  query: PropTypes.Validator<object>;
  subscription: PropTypes.Validator<object>;
}

export default function FeedSubscription({ subscription, query }: Props) {
  return (
    <Subscription
      subscription={subscription}
      onSubscriptionData={options => onReceivedData({ ...options, query })}
    >
      {() => null}
    </Subscription>
  );
}
