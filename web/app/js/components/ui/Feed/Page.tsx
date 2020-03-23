import React, { useState, PropsWithChildren, useContext } from "react";
import { makeStyles } from "@material-ui/core/styles";
import CreateBookmark from "../Panel/CreateBookmark";
import { AddButton } from "../Fab";
import Grid from "../Grid";
import { LayoutContext, MessageContext } from "../../context";
import { FEED_SM_WIDTH, FEED_LG_WIDTH } from "../../../constant/layout";

const useStyles = makeStyles(({ breakpoints }) => ({
  root: {
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
  },
  content: {
    [breakpoints.up("sm")]: {
      width: FEED_SM_WIDTH,
    },
    [breakpoints.up("md")]: {
      width: FEED_LG_WIDTH,
    },
  },
}));

export default function FeedPage({
  children,
}: PropsWithChildren<{}>): JSX.Element {
  const classes = useStyles();
  const setMessageInfo = useContext(MessageContext);
  const setIsContained = useContext(LayoutContext);
  const [isPanelOpen, setIsPanelOpen] = useState(false);

  return (
    <>
      <Grid className={classes.content}>{children}</Grid>
      <AddButton
        onClick={() => {
          setIsContained(true);
          setIsPanelOpen(true);
        }}
      />
      <CreateBookmark
        isOpen={isPanelOpen}
        setIsPanelOpen={(isOpen) => {
          setIsContained(isOpen);
          setIsPanelOpen(isOpen);
        }}
        onBookmarkCreated={() => {
          setIsContained(false);
          setIsPanelOpen(false);
          setMessageInfo({
            message: "Nice one! The bookmark was added",
          });
        }}
      />
    </>
  );
}
