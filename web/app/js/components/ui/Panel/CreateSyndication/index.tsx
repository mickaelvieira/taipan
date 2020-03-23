import React from "react";
import { Source } from "../../../../types/syndication";
import Panel from "..";
import Form from "./Form";

interface Props {
  isOpen: boolean;
  toggleDialog: (status: boolean) => void;
  onSyndicationCreated: (source: Source) => void;
}

export default function AddSubscription({
  isOpen,
  toggleDialog,
  onSyndicationCreated,
}: Props): JSX.Element {
  return (
    <Panel isOpen={isOpen} prev={() => toggleDialog(false)} title="Add a feed">
      <Form onCreated={onSyndicationCreated} />
    </Panel>
  );
}
