import React, { PropsWithChildren } from "react";
import Paper from "@material-ui/core/Paper";
import { makeStyles } from "@material-ui/core/styles";

const useStyle = makeStyles({
  message: {
    padding: 24
  }
});

interface Props {
  className?: string;
}

export default function EmptyFeed({
  className,
  children
}: PropsWithChildren<Props>): JSX.Element {
  const classes = useStyle();
  return (
    <Paper className={`${classes.message} ${className ? className : ""}`}>
      {children}
    </Paper>
  );
}
