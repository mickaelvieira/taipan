import React, { useState, PropsWithChildren } from "react";
import CreateBookmark from "./Panel/CreateBookmark";
import { AddButton } from "../Fab";
import { LayoutRenderProps } from "../../Layout/Layout";
import Grid from "../Grid";

export default function FeedPage({
  children,
  setMessageInfo,
  setIsContained
}: PropsWithChildren<LayoutRenderProps>): JSX.Element {
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
