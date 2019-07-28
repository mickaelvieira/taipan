import React, { PropsWithChildren } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Slide from "@material-ui/core/Slide";
import Paper from "@material-ui/core/Paper";

const useStyles = makeStyles(() => ({
  paper: {
    zIndex: 10000,
    position: "fixed",
    width: "100vw",
    height: "100vh"
  }
}));

interface Props {
  isOpen: boolean;
  prev: () => void;
}

export default function AddForm({
  isOpen,
  children
}: PropsWithChildren<Props>): JSX.Element {
  const classes = useStyles();

  return (
    <Slide in={isOpen} direction="up" mountOnEnter unmountOnExit>
      <Paper className={classes.paper} elevation={0} square>
        {children}
      </Paper>
    </Slide>
  );
}
