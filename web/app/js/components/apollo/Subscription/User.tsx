import React from "react";
import { Subscription as SubscriptionBase } from "react-apollo";
import subscription from "../graphql/subscription/user.graphql";
import { UserEvent } from "../../../types/users";

interface Data {
  userChanged: UserEvent;
}

export { subscription };

class Subscription extends SubscriptionBase<Data, {}> {}

export default function UserSubscription(): JSX.Element {
  return <Subscription subscription={subscription} />;
}
