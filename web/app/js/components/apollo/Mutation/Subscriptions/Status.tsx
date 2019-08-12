import { Subscription } from "../../../../types/subscription";
import subscribeMutation from "../../graphql/mutation/subscriptions/subscribe.graphql";
import unsubscribeMutation from "../../graphql/mutation/subscriptions/unsubscribe.graphql";

export interface Data {
  syndication: {
    subscribe?: Subscription;
    unsubscribe?: Subscription;
  };
}

export interface Variables {
  url: URL;
}

export { subscribeMutation, unsubscribeMutation };
