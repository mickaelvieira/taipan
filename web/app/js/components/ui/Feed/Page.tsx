import React, { useState, PropsWithChildren, useContext } from "react";
import CreateBookmark from "./Panel/CreateBookmark";
import { AddButton } from "../Fab";
import Grid from "../Grid";
import { LayoutContext, MessageContext } from "../../context";

export default function FeedPage({
  children
}: PropsWithChildren<{}>): JSX.Element {
  const setMessageInfo = useContext(MessageContext);
  const setIsContained = useContext(LayoutContext);
  const [isPanelOpen, setIsPanelOpen] = useState(false);

  return (
    <>
      <Grid>{children}</Grid>
      <AddButton
        onClick={() => {
          setIsContained(true);
          setIsPanelOpen(true);
        }}
      />
      <CreateBookmark
        isOpen={isPanelOpen}
        setIsPanelOpen={isOpen => {
          setIsContained(isOpen);
          setIsPanelOpen(isOpen);
        }}
        onBookmarkCreated={() => {
          setIsContained(false);
          setIsPanelOpen(false);
          setMessageInfo({
            message: "Nice one! The bookmark was added"
          });
        }}
      />
    </>
  );
}
