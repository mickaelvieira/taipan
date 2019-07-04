import React, { useState, PropsWithChildren, useContext } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Fab from "@material-ui/core/Fab";
import AddIcon from "@material-ui/icons/Add";
import Grid from "@material-ui/core/Grid";
import AddBookmark from "../ui/Feed/Modal/AddBookmark";
import { MessageContext } from "../context";
import { LG_WIDTH, SM_WIDTH } from "../../constant/feed";
import MainLayout from "./Layout";

const useStyles = makeStyles(({ breakpoints, palette, spacing }) => ({
  root: {
    width: "100%",
    maxWidth: SM_WIDTH,
    [breakpoints.up("md")]: {
      maxWidth: LG_WIDTH
    }
  },
  content: {
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
    color: palette.text.secondary,
    paddingTop: 70
  },
  fab: {
    margin: spacing(1),
    position: "fixed",
    bottom: spacing(2),
    right: spacing(2),
    backgroundColor: palette.secondary.main,
    "&:hover": {
      backgroundColor: palette.secondary.light
    }
  }
}));

export default function LayoutFeed({
  children
}: PropsWithChildren<{}>): JSX.Element {
  const classes = useStyles();
  const setInfo = useContext(MessageContext);
  const [isModalOpen, setModalStatus] = useState(false);

  return (
    <MainLayout>
      <Grid item xs={12} className={classes.root}>
        <div className={classes.content}>{children}</div>
        <Fab
          color="primary"
          size="large"
          aria-label="Add"
          className={classes.fab}
          onClick={() => setModalStatus(true)}
        >
          <AddIcon />
        </Fab>
        <AddBookmark
          isOpen={isModalOpen}
          toggleDialog={setModalStatus}
          onBookmarkCreated={() => {
            setInfo("Nice one! The bookmark was added");
            setModalStatus(false);
          }}
        />
      </Grid>
    </MainLayout>
  );
}
