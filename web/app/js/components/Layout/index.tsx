import React, { useState, PropsWithChildren } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Fab from "@material-ui/core/Fab";
import AddIcon from "@material-ui/icons/Add";
import Grid from "@material-ui/core/Grid";
import Header from "./Header";
import Sidebar from "./Navigation/Sidebar";
import AddForm from "../AddForm";
import useConnectionStatus from "../../hooks/connection-status";
import { SnackbarInfo } from "../ui/Snackbar";
import { MessageContext } from "../context";

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

export default function Layout({
  children
}: PropsWithChildren<{}>): JSX.Element {
  const classes = useStyles();
  const [info, setInfo] = useState("");
  const [isOpen, setIsOpen] = useState(false);
  const [isFormBookmarkOpen, setFormBookmarkStatus] = useState(false);
  const isOnline = useConnectionStatus();

  return (
    <>
      <Sidebar isOpen={isOpen} toggleDrawer={setIsOpen} />
      <Header toggleDrawer={setIsOpen} />
      <Grid container className={classes.root}>
        <Grid item xs={12} className={classes.inner}>
          <MessageContext.Provider value={setInfo}>
            <div className={classes.content}>{children}</div>
          </MessageContext.Provider>
          <Fab
            color="primary"
            size="large"
            aria-label="Add"
            className={classes.fab}
            onClick={() => setFormBookmarkStatus(true)}
          >
            <AddIcon />
          </Fab>
        </Grid>
      </Grid>
      <AddForm
        isOpen={isFormBookmarkOpen}
        toggleDialog={setFormBookmarkStatus}
        onSyndicationSourceCreated={() => {
          setInfo("Nice one! The feed was added");
          setFormBookmarkStatus(false);
        }}
        onBookmarkCreated={() => {
          setInfo("Nice one! The bookmark was added");
          setFormBookmarkStatus(false);
        }}
      />
      <SnackbarInfo open={!isOnline} info="You are offline" />
      <SnackbarInfo
        onClose={() => setInfo("")}
        autoHideDuration={3000}
        open={info !== ""}
        info={info}
      />
    </>
  );
}
