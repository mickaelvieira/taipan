import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import Paper from "@material-ui/core/Paper";
import Terms from "./Terms";
import { SearchType } from "../../../types/search";
import Emoji from "../../ui/Emoji";

const useStyles = makeStyles(({ spacing }) => ({
  message: {
    padding: spacing(2)
  },
  type: {
    fontWeight: 500
  }
}));

interface Props {
  type: SearchType;
  terms: string[];
}

export default function NoResults({ type, terms }: Props): JSX.Element {
  const classes = useStyles();
  return (
    <Paper className={classes.message}>
      Sorry, I could not find any <span className={classes.type}>{type}s</span>{" "}
      matching{" "}
      {terms.length > 0 ? (
        <Terms terms={terms} />
      ) : (
        <span>
          nothing
          <Emoji emoji=":zany_face:" />
        </span>
      )}
    </Paper>
  );
}
