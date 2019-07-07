import React, { useState, PropsWithChildren } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Grid from "@material-ui/core/Grid";
import Header from "./Header";
import Sidebar from "./Navigation/Sidebar";
import useConnectionStatus from "../../hooks/connection-status";
import { SnackbarInfo } from "../ui/Snackbar";
import { MessageContext } from "../context";

const useStyles = makeStyles(() => ({
  root: {
    display: "flex"
  }
}));

export default function Layout({
  children
}: PropsWithChildren<{}>): JSX.Element {
  const classes = useStyles();
  const [info, setInfo] = useState("");
  const [isOpen, setIsOpen] = useState(false);
  const isOnline = useConnectionStatus();

  return (
    <div className={classes.root}>
      <Sidebar isOpen={isOpen} toggleDrawer={setIsOpen} />
      <Header toggleDrawer={setIsOpen} />
      <Grid container>
        <MessageContext.Provider value={setInfo}>
          {children}
        </MessageContext.Provider>
      </Grid>
      <SnackbarInfo open={!isOnline} info="You are offline" />
      <SnackbarInfo
        onClose={() => setInfo("")}
        autoHideDuration={3000}
        open={info !== ""}
        info={info}
      />
    </div>
  );
}
