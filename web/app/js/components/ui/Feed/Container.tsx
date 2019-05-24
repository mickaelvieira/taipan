import React, { PropsWithChildren } from "react";
import { makeStyles } from "@material-ui/core/styles";

const useStyles = makeStyles({
  container: {
    marginBottom: 60,
    width: "100%",
    minHeight: "100vh"
  }
});

export default function FeedContainer({ children }: PropsWithChildren<{}>) {
  const classes = useStyles();
  return <section className={classes.container}>{children}</section>;
}
