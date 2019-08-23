import React from "react";
import { Subscription } from "../../../../types/subscription";
import Panel from "..";
import Form from "./Form";

interface Props {
  isOpen: boolean;
  toggleDialog: (status: boolean) => void;
  onSubscriptionCreated: (subscription: Subscription) => void;
}

export default function AddSubscription({
  isOpen,
  toggleDialog,
  onSubscriptionCreated
}: Props): JSX.Element {
  return (
    <Panel
      isOpen={isOpen}
      prev={() => toggleDialog(false)}
      title="Subscribe to feed"
    >
      <Form onCreated={onSubscriptionCreated} />
    </Panel>
  );
}
