import React from "react";
import Paper from "@material-ui/core/Paper";
import { makeStyles } from "@material-ui/core/styles";

const useStyle = makeStyles({
  message: {
    padding: 24
  }
});

interface Props {
  message: string;
}

export default function EmptyFeed({ message }: Props): JSX.Element {
  const classes = useStyle();
  return <Paper className={classes.message}>{message}</Paper>;
}
