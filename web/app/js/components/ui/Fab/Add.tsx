import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import Fab from "@material-ui/core/Fab";
import AddIcon from "@material-ui/icons/Add";

const useStyles = makeStyles(({ palette, spacing }) => ({
  fab: {
    margin: spacing(1),
    position: "fixed",
    bottom: spacing(2),
    right: spacing(2),
    backgroundColor: palette.secondary.main,
    "&:hover": {
      backgroundColor: palette.secondary.light
    }
  }
}));

interface Props {
  onClick: () => void;
}

export default function FabAddButton({ onClick }: Props): JSX.Element {
  const classes = useStyles();
  return (
    <Fab
      color="primary"
      size="large"
      aria-label="Add"
      className={classes.fab}
      onClick={onClick}
    >
      <AddIcon />
    </Fab>
  );
}
