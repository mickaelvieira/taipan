import { DataProxy } from "apollo-cache";
import { Mutation } from "react-apollo";
import { Subscription } from "../../../../types/subscription";
import mutation from "../../graphql/mutation/subscriptions/subscription.graphql";
import { query, Data as QueryData } from "../../Query/Subscriptions";
import { addSubscription } from "../../helpers/subscriptions";

interface Data {
  subscriptions: {
    subscription: Subscription;
  };
}

interface Variables {
  url: string;
}

class SubscriptionMutation extends Mutation<Data, Variables> {
  static defaultProps = {
    mutation,
    update: (cache: DataProxy, { data }: { data: Data }) => {
      const { subscription } = data.subscriptions;
      const prev = cache.readQuery({ query }) as QueryData;
      const result = addSubscription(
        prev.subscriptions.subscriptions,
        subscription
      );
      cache.writeQuery({
        query,
        data: {
          syndication: {
            ...prev.subscriptions,
            subscriptions: result
          }
        }
      });
    }
  };
}

export { mutation };

export default SubscriptionMutation;
