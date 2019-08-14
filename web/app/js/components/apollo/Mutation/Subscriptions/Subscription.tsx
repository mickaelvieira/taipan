import { Subscription } from "../../../../types/subscription";
import mutation from "../../graphql/mutation/subscriptions/subscription.graphql";

export interface Data {
  subscriptions: {
    subscription: Subscription;
  };
}

export interface Variables {
  url: string;
}

export { mutation };
