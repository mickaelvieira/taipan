import React, { PropsWithChildren } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Card from "@material-ui/core/Card";

const useStyles = makeStyles(({ breakpoints }) => ({
  card: {
    marginBottom: 24,
    display: "flex",
    flexDirection: "column",
    borderRadius: 0,
    [breakpoints.up("sm")]: {
      borderRadius: 4
    }
  }
}));

export default function Item({ children }: PropsWithChildren<{}>) {
  const classes = useStyles();
  return <Card className={classes.card}>{children}</Card>;
}
