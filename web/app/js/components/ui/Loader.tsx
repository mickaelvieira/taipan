import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import CircularProgress from "@material-ui/core/CircularProgress";

const useStyles = makeStyles(({ spacing }) => ({
  container: {
    flex: 1,
    display: "flex",
    justifyContent: "center",
    alignItems: "center",
  },
  progress: {
    margin: spacing(2),
  },
}));

export default function Loader(): JSX.Element {
  const classes = useStyles();
  return (
    <div className={classes.container}>
      <CircularProgress className={classes.progress} />
    </div>
  );
}
