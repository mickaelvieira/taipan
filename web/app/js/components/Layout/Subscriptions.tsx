import React, { useState, PropsWithChildren } from "react";
import AddSubscriptionModal from "../ui/Subscriptions/Modal/AddSubscription";
import { AddButton } from "../ui/Fab";
import MainLayout from "./Layout";
import MainContent from "./Content";

export default function SubscriptionLayout({
  children
}: PropsWithChildren<{}>): JSX.Element {
  const [isModalOpen, setModalStatus] = useState(false);

  return (
    <MainLayout>
      {({ setInfoMessage }) => (
        <>
          <MainContent>{children}</MainContent>
          <AddButton onClick={() => setModalStatus(true)} />
          <AddSubscriptionModal
            isOpen={isModalOpen}
            toggleDialog={setModalStatus}
            onSubscriptionCreated={() => {
              setInfoMessage("Nice one! The feed was added");
              setModalStatus(false);
            }}
          />
        </>
      )}
    </MainLayout>
  );
}
