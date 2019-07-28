import React, { useState, PropsWithChildren } from "react";
import AddSubscriptionModal from "../ui/Subscriptions/Panel/AddSubscription";
import { AddButton } from "../ui/Fab";
import MainLayout from "./Layout";
import MainContent from "./Content";

export default function SubscriptionLayout({
  children
}: PropsWithChildren<{}>): JSX.Element {
  const [isModalOpen, setModalStatus] = useState(false);

  return (
    <MainLayout>
      {({ setMessageInfo, setIsContained }) => (
        <>
          <MainContent>{children}</MainContent>
          <AddButton
            onClick={() => {
              setIsContained(true);
              setModalStatus(true);
            }}
          />
          <AddSubscriptionModal
            isOpen={isModalOpen}
            toggleDialog={status => {
              setIsContained(status);
              setModalStatus(status);
            }}
            onSubscriptionCreated={() => {
              setMessageInfo({ message: "Nice one! The feed was added" });
              setModalStatus(false);
            }}
          />
        </>
      )}
    </MainLayout>
  );
}
