import { Mutation } from "react-apollo";
import { Subscription } from "../../../../types/subscription";
import mutation from "../../graphql/mutation/subscriptions/subscription.graphql";

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
    mutation
  };
}

export { mutation };

export default SubscriptionMutation;
