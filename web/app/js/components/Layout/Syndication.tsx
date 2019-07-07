import React, { useState, PropsWithChildren, useContext } from "react";
import AddSourceModal from "../ui/Syndication/Modal/AddSource";
import { AddButton } from "../ui/Fab";
import { MessageContext } from "../context";
import MainLayout from "./Layout";
import MainContent from "./Content";

export default function LayoutSyndication({
  children
}: PropsWithChildren<{}>): JSX.Element {
  const setInfo = useContext(MessageContext);
  const [isModalOpen, setModalStatus] = useState(false);

  return (
    <MainLayout>
      <MainContent>{children}</MainContent>
      <AddButton onClick={() => setModalStatus(true)} />
      <AddSourceModal
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
