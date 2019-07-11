import React, { useState, PropsWithChildren, useContext } from "react";
import AddBookmarkModal from "../ui/Feed/Modal/AddBookmark";
import { AddButton } from "../ui/Fab";
import { MessageContext } from "../context";
import MainLayout from "./Layout";
import MainContent from "./Content";

export default function LayoutFeed({
  children
}: PropsWithChildren<{}>): JSX.Element {
  const setInfo = useContext(MessageContext);
  const [isModalOpen, setModalStatus] = useState(false);

  return (
    <MainLayout>
      <MainContent>{children}</MainContent>
      <AddButton onClick={() => setModalStatus(true)} />
      <AddBookmarkModal
        isOpen={isModalOpen}
        toggleDialog={setModalStatus}
        onBookmarkCreated={() => {
          setInfo("Nice one! The bookmark was added");
          setModalStatus(false);
        }}
      />
    </MainLayout>
  );
}
