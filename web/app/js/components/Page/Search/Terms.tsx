import React from "react";
import { makeStyles } from "@material-ui/core/styles";

const useStyles = makeStyles(({ palette }) => ({
  term: {
    display: "inline-block",
    padding: "0 4px",
    margin: "0 2px",
    color: palette.common.white,
    backgroundColor: palette.primary.main
  }
}));

interface Props {
  terms: string[];
}

export default function Terms({ terms }: Props): JSX.Element {
  const classes = useStyles();
  return (
    <>
      {terms.map(term => (
        <span className={classes.term} key={term}>
          {term}
        </span>
      ))}
    </>
  );
}
