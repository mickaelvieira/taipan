import { Subscription, OnSubscriptionDataOptions } from "react-apollo";
import subscription from "../../../services/apollo/subscription/favorites.graphql";
import {
  queryFavorites as query,
  addItemFromFeedResults,
  getDataKey
} from "../Query/Feed";
import { Bookmark } from "../../../types/bookmark";

export interface Data {
  LatestFavorite: {
    id: string;
    bookmark: Bookmark;
  };
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
        const item = subscriptionData.data.LatestFavorite.bookmark;
        const data = client.readQuery({ query });
        if (data) {
          const key = getDataKey(data);
          if (key) {
            const result = addItemFromFeedResults(data[key], item);
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
