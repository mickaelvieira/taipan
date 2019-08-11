import React, { PropsWithChildren } from "react";
import { makeStyles } from "@material-ui/core/styles";

const useStyles = makeStyles(() => ({
  root: {
    display: "flex",
    flexDirection: "column",
    height: "100vh",
    justifyContent: "center",
    alignItems: "center"
  },
  contained: {
    overflow: "hidden"
  }
}));

export default function OutLayout({
  children
}: PropsWithChildren<{}>): JSX.Element {
  const classes = useStyles();

  return <div className={classes.root}>{children}</div>;
}
