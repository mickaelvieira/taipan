import React, { PropsWithChildren } from "react";
import { makeStyles } from "@material-ui/core/styles";

const useStyles = makeStyles({
  footer: {
    display: "flex",
    flexDirection: "row",
    justifyContent: "space-between"
  }
});

export default function ItemFooter({ children }: PropsWithChildren<{}>) {
  const classes = useStyles();
  return <footer className={classes.footer}>{children}</footer>;
}
