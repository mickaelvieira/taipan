import React, { PropsWithChildren } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Paper from "@material-ui/core/Paper";

const useStyles = makeStyles(() => ({
  message: {
    padding: 24,
  },
}));

export default function Empty({
  children,
}: PropsWithChildren<{}>): JSX.Element {
  const classes = useStyles();
  return <Paper className={classes.message}>{children}</Paper>;
}
