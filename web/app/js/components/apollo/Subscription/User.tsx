import React, { useContext } from "react";
import { Subscription as SubscriptionBase } from "react-apollo";
import subscription from "../graphql/subscription/user.graphql";
import { Event } from "../../../types/subscription";
import { User, UserEvent } from "../../../types/users";
import { ClientContext } from "../../context";

function hasReceivedEvent(data: Data | undefined): [boolean, UserEvent | null] {
  let isReceived = false;
  let event: UserEvent | null = null;

  if (data && data.user) {
    event = data.user;
    isReceived = true;
  }

  return [isReceived, event];
}

function isEmitter(event: Event | null, clientId: string): boolean {
  if (!event) {
    return false;
  }
  return event.emitter === clientId;
}

interface Data {
  user: UserEvent;
}

export { subscription };

class Subscription extends SubscriptionBase<Data, {}> {}

interface Props {
  update: (user: User) => void;
}

export default function UserSubscription({ update }: Props): JSX.Element {
  const clientId = useContext(ClientContext);
  return (
    <Subscription
      subscription={subscription}
      onSubscriptionData={({ subscriptionData }) => {
        const [isReceived, event] = hasReceivedEvent(subscriptionData.data);
        console.log(event);
        console.log(clientId);
        if (isReceived && !isEmitter(event, clientId)) {
          const { item } = event as UserEvent;
          update(item);
        }
      }}
    />
  );
}
