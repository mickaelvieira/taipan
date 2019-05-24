import React, { useState, PropsWithChildren } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Fab from "@material-ui/core/Fab";
import AddIcon from "@material-ui/icons/Add";
import Grid from "@material-ui/core/Grid";
import Header from "./Header";
import Sidebar from "./Navigation/Sidebar";
import AddBookmark from "../AddBookmark";
import useConnectionStatus from "../../hooks/connection-status";
import { SnackbarInfo } from "../ui/Snackbar";

const useStyles = makeStyles(({ palette, spacing }) => ({
  root: {
    display: "flex",
    flexDirection: "column",
    alignItems: "center"
  },
  inner: {
    width: "100%",
    maxWidth: 600
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
  },
  message: {
    display: "flex",
    alignItems: "center"
  }
}));

export default function Layout({ children }: PropsWithChildren<{}>) {
  const classes = useStyles();
  const [isOpen, setIsOpen] = useState(false);
  const [isFormBookmarkOpen, setFormBookmarkStatus] = useState(false);
  const isOnline = useConnectionStatus();

  return (
    <>
      <Sidebar isOpen={isOpen} toggleDrawer={setIsOpen} />
      <Header toggleDrawer={setIsOpen} />
      <Grid container className={classes.root}>
        <Grid item xs={12} className={classes.inner}>
          <div className={classes.content}>{children}</div>
          <Fab
            color="primary"
            size="small"
            aria-label="Add"
            className={classes.fab}
            onClick={() => setFormBookmarkStatus(true)}
          >
            <AddIcon />
          </Fab>
        </Grid>
      </Grid>
      <AddBookmark
        isOpen={isFormBookmarkOpen}
        toggleDialog={setFormBookmarkStatus}
        onBookmarkCreated={() => setFormBookmarkStatus(false)}
      />
      <SnackbarInfo isOpen={!isOnline} message="You are offline" />
    </>
  );
}
