import React, { PropsWithChildren } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Paper from "@material-ui/core/Paper";

const useStyles = makeStyles(() => ({
  paper: {
    display: "flex",
    flexDirection: "column"
  },
  hidden: {
    display: "none"
  }
}));

interface Props {
  isVisible: boolean;
  className?: string;
}

export default function TabPanel({
  children,
  isVisible,
  className
}: PropsWithChildren<Props>): JSX.Element | null {
  const classes = useStyles();
  let styles = [classes.paper];
  if (!isVisible) {
    styles.push(classes.hidden);
  }
  if (className) {
    styles.push(className);
  }

  return (
    <Paper elevation={0} className={`${styles.join(" ")}`}>
      {children}
    </Paper>
  );
}
