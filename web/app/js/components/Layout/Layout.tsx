import React, { useState, useEffect } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Grid from "@material-ui/core/Grid";
import Header from "./Header";
import Sidebar from "./Navigation/Sidebar";
import useConnectionStatus from "../../hooks/useConnectionStatus";
import { SnackbarInfo } from "../ui/Snackbar";
import { MessageContext } from "../context";
import { MessageInfo } from "../../types";

const useStyles = makeStyles(() => ({
  root: {
    display: "flex",
    height: "100vh"
  },
  contained: {
    overflow: "hidden"
  }
}));

export interface LayoutRenderProps {
  setMessageInfo: (message: MessageInfo | null) => void;
  setIsContained: (contained: boolean) => void;
  setIsSidebarOpen: (open: boolean) => void;
}

interface Props {
  children: (props: LayoutRenderProps) => JSX.Element;
}

export default function Layout({ children }: Props): JSX.Element {
  const classes = useStyles();
  const [info, setMessageInfo] = useState<MessageInfo | null>(null);
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
        <MessageContext.Provider value={setMessageInfo}>
          {children({
            setMessageInfo,
            setIsContained,
            setIsSidebarOpen
          })}
        </MessageContext.Provider>
      </Grid>
      <SnackbarInfo open={!isOnline} info={{ message: "You are offline" }} />
      <SnackbarInfo
        onClose={() => setMessageInfo(null)}
        forceClose={() => setMessageInfo(null)}
        autoHideDuration={3000}
        open={info !== null}
        info={info}
      />
    </div>
  );
}
