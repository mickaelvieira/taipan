import React, { PropsWithChildren, useState, useEffect } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Grid from "@material-ui/core/Grid";
import UserContextProvider from "../context/provider/User";
import FeedContextProvider from "../context/provider/Feeds";
import Header from "./Header";
import Sidebar from "./Navigation/Sidebar";
import useConnectionStatus from "../../hooks/useConnectionStatus";
import { SnackbarInfo } from "../ui/Snackbar";
import { MessageContext, LayoutContext } from "../context";
import { MessageInfo } from "../../types";
import { User } from "../../types/users";

const useStyles = makeStyles(() => ({
  root: {
    display: "flex",
    height: "100vh"
  },
  contained: {
    overflow: "hidden"
  }
}));

interface Props {
  user: User | null;
}

export default function AppLayout({
  user,
  children
}: PropsWithChildren<Props>): JSX.Element | null {
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

  return !user ? null : (
    <UserContextProvider user={user}>
      <FeedContextProvider>
        <div className={classes.root}>
          <Sidebar isOpen={isSideOpen} toggleDrawer={setIsSidebarOpen} />
          <Header toggleDrawer={setIsSidebarOpen} />
          <Grid container>
            <LayoutContext.Provider value={setIsContained}>
              <MessageContext.Provider value={setMessageInfo}>
                {children}
              </MessageContext.Provider>
            </LayoutContext.Provider>
          </Grid>
          <SnackbarInfo
            open={!isOnline}
            info={{ message: "You are offline" }}
          />
          <SnackbarInfo
            onClose={() => setMessageInfo(null)}
            forceClose={() => setMessageInfo(null)}
            autoHideDuration={3000}
            open={info !== null}
            info={info}
          />
        </div>
      </FeedContextProvider>
    </UserContextProvider>
  );
}
