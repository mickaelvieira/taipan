import { Subscription, OnSubscriptionDataOptions } from "react-apollo";
import subscription from "../../../services/apollo/subscription/favorites.graphql";
import { queryFavorites as query, getDataKey } from "../Query/Feed";
import { FeedEventData, FeedEvent } from "../../../types/feed";
import { feedResultsAction } from "../helpers/feed";
import { hasReceivedEvent } from "../helpers/data";

const variables = {};

export { subscription, variables };

class LatestFavoriteSubscription extends Subscription<FeedEventData, {}> {
  static defaultProps = {
    subscription,
    onSubscriptionData: ({
      client,
      subscriptionData
    }: OnSubscriptionDataOptions<FeedEventData>) => {
      const [isReceived, event] = hasReceivedEvent(subscriptionData.data);
      if (isReceived) {
        console.log(event);
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
  };
}

export default LatestFavoriteSubscription;
