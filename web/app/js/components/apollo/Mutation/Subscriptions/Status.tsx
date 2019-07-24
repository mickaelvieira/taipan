import { Mutation } from "react-apollo";
import { Subscription } from "../../../../types/subscription";
import subscribeMutation from "../../graphql/mutation/subscriptions/subscribe.graphql";
import unsubscribeMutation from "../../graphql/mutation/subscriptions/unsubscribe.graphql";

interface Data {
  syndication: {
    subscribe?: Subscription;
    unsubscribe?: Subscription;
  };
}

interface Variables {
  url: string;
}

class ChangeStatusMutation extends Mutation<Data, Variables> {
  static defaultProps = {};
}

export { subscribeMutation, unsubscribeMutation };

export default ChangeStatusMutation;
