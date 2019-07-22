import React, { useState, useEffect } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Grid from "@material-ui/core/Grid";
import Header from "./Header";
import Sidebar from "./Navigation/Sidebar";
import useConnectionStatus from "../../hooks/connection-status";
import { SnackbarInfo } from "../ui/Snackbar";
import { MessageContext } from "../context";

const useStyles = makeStyles(() => ({
  root: {
    display: "flex",
    height: "100vh"
  },
  contained: {
    overflow: "hidden"
  }
}));

interface RenderProps {
  setInfoMessage: (message: string) => void;
  setIsContained: (contained: boolean) => void;
  setIsSidebarOpen: (open: boolean) => void;
}

interface Props {
  children: (props: RenderProps) => JSX.Element;
}

export default function Layout({ children }: Props): JSX.Element {
  const classes = useStyles();
  const [info, setInfoMessage] = useState("");
  const [isSideOpen, setIsSidebarOpen] = useState(false);
  const [isContained, setIsContained] = useState(false);
  const isOnline = useConnectionStatus();
  const body = document.querySelector("body");

  useEffect(() => {
    const overflow = isContained ? "hidden" : "initial";
    if (body) {
      body.style.overflow = overflow;
    }
  }, [body, isContained]);

  return (
    <div className={classes.root}>
      <Sidebar isOpen={isSideOpen} toggleDrawer={setIsSidebarOpen} />
      <Header toggleDrawer={setIsSidebarOpen} />
      <Grid container>
        <MessageContext.Provider value={setInfoMessage}>
          {children({
            setInfoMessage,
            setIsContained,
            setIsSidebarOpen
          })}
        </MessageContext.Provider>
      </Grid>
      <SnackbarInfo open={!isOnline} info="You are offline" />
      <SnackbarInfo
        onClose={() => setInfoMessage("")}
        autoHideDuration={3000}
        open={info !== ""}
        info={info}
      />
    </div>
  );
}
