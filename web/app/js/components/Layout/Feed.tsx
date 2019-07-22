import React, { useState, PropsWithChildren } from "react";
import AddBookmarkPanel from "../ui/Feed/Panel/CreateBookmark";
import { AddButton } from "../ui/Fab";
import MainLayout from "./Layout";
import MainContent from "./Content";

export default function FeedLayout({
  children
}: PropsWithChildren<{}>): JSX.Element {
  const [isPanelOpen, setIsPanelOpen] = useState(false);

  return (
    <MainLayout>
      {({ setInfoMessage, setIsContained }) => (
        <>
          <MainContent>{children}</MainContent>
          <AddButton
            onClick={() => {
              setIsContained(true);
              setIsPanelOpen(true);
            }}
          />
          <AddBookmarkPanel
            isOpen={isPanelOpen}
            setIsPanelOpen={isOpen => {
              setIsContained(isOpen);
              setIsPanelOpen(isOpen);
            }}
            onBookmarkCreated={() => {
              setIsContained(false);
              setIsPanelOpen(false);
              setInfoMessage("Nice one! The bookmark was added");
            }}
          />
        </>
      )}
    </MainLayout>
  );
}
