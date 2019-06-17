import { Subscription, OnSubscriptionDataOptions } from "react-apollo";
import subscription from "../../../services/apollo/subscription/reading-list.graphql";
import { FeedEvent } from "../../../types/feed";
import { queryReadingList as query, getDataKey } from "../Query/Feed";
import { feedResultsAction } from "../helpers/feed";

export interface Data {
  LatestReadingList: FeedEvent;
}

const variables = {};

export { query, variables };

class ReadingListSubscription extends Subscription<Data, {}> {
  static defaultProps = {
    subscription,
    onSubscriptionData: ({
      client,
      subscriptionData
    }: OnSubscriptionDataOptions<Data>) => {
      if (subscriptionData.data) {
        console.log(subscriptionData.data.LatestReadingList);
        const { bookmark, action } = subscriptionData.data.LatestReadingList;
        const updateCache = feedResultsAction[action];
        const data = client.readQuery({ query });
        if (data) {
          const key = getDataKey(data);
          if (key) {
            const result = updateCache(data[key], bookmark);
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

export default ReadingListSubscription;
