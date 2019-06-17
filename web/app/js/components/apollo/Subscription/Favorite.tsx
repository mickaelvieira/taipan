import { Subscription, OnSubscriptionDataOptions } from "react-apollo";
import subscription from "../../../services/apollo/subscription/favorites.graphql";
import { queryFavorites as query, getDataKey } from "../Query/Feed";
import { BookmarkEvent } from "../../../types/feed";
import { feedResultsAction } from "../helpers/feed";

export interface Data {
  LatestFavorite: BookmarkEvent;
}

const variables = {};

export { subscription, variables };

class LatestFavoriteSubscription extends Subscription<Data, {}> {
  static defaultProps = {
    subscription,
    onSubscriptionData: ({
      client,
      subscriptionData
    }: OnSubscriptionDataOptions<Data>) => {
      if (subscriptionData.data) {
        console.log(subscriptionData.data.LatestFavorite);
        const { bookmark, action } = subscriptionData.data.LatestFavorite;
        const data = client.readQuery({ query });
        const updateResult = feedResultsAction[action];
        if (data) {
          const key = getDataKey(data);
          if (key) {
            const result = updateResult(data[key], bookmark);
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
