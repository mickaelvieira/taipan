import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import Paper from "@material-ui/core/Paper";
import Terms from "./Terms";

const useStyles = makeStyles(({ spacing }) => ({
  message: {
    padding: spacing(2)
  }
}));

interface Props {
  terms: string[];
}

export default function NoResults({ terms }: Props): JSX.Element {
  const classes = useStyles();
  return (
    <Paper className={classes.message}>
      Sorry, we could not find any sources matching <Terms terms={terms} />
    </Paper>
  );
}
