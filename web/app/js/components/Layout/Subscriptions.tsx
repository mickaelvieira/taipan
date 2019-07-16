import React, { useState, PropsWithChildren, useContext } from "react";
import AddSubscriptionModal from "../ui/Subscriptions/Modal/AddSubscription";
import { AddButton } from "../ui/Fab";
import { MessageContext } from "../context";
import MainLayout from "./Layout";
import MainContent from "./Content";

export default function LayoutSubscription({
  children
}: PropsWithChildren<{}>): JSX.Element {
  const setInfo = useContext(MessageContext);
  const [isModalOpen, setModalStatus] = useState(false);

  return (
    <MainLayout>
      <MainContent>{children}</MainContent>
      <AddButton onClick={() => setModalStatus(true)} />
      <AddSubscriptionModal
        isOpen={isModalOpen}
        toggleDialog={setModalStatus}
        onSyndicationSourceCreated={() => {
          setInfo("Nice one! The feed was added");
          setModalStatus(false);
        }}
      />
    </MainLayout>
  );
}
