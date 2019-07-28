import React from "react";
import { Subscription } from "../../../../../types/subscription";
import Panel from "../../../Panel";
import FormSubscription from "./FormSubscription";

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
    <Panel isOpen={isOpen} prev={() => toggleDialog(false)} title="Add a feed">
      <FormSubscription onSubscriptionCreated={onSubscriptionCreated} />
    </Panel>
  );
}
