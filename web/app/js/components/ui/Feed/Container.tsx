import React, { PropsWithChildren } from "react";
import { makeStyles } from "@material-ui/core/styles";

const useStyles = makeStyles({
  container: {
    display: "flex",
    flexDirection: "column",
    marginBottom: 60,
    minHeight: "100vh"
  }
});

export default function FeedContainer({
  children
}: PropsWithChildren<{}>): JSX.Element {
  const classes = useStyles();
  return <section className={classes.container}>{children}</section>;
}
