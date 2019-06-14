import { Subscription, OnSubscriptionDataOptions } from "react-apollo";
import subscription from "../../../services/apollo/subscription/reading-list.graphql";
import { Bookmark } from "../../../types/bookmark";
import {
  queryReadingList as query,
  addItemFromFeedResults,
  getDataKey
} from "../Query/Feed";

export interface Data {
  LatestReadingList: {
    id: string;
    bookmark: Bookmark;
  };
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
        const item = subscriptionData.data.LatestReadingList.bookmark;
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

export default ReadingListSubscription;
