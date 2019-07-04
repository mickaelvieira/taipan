import React, { PropsWithChildren } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Grid from "@material-ui/core/Grid";
import MainLayout from "./Layout";

const useStyles = makeStyles(({ palette }) => ({
  root: {
    width: "100%"
  },
  content: {
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
    color: palette.text.secondary,
    paddingTop: 70
  }
}));

export default function LayoutSyndication({
  children
}: PropsWithChildren<{}>): JSX.Element {
  const classes = useStyles();

  return (
    <MainLayout>
      <Grid item xs={12} className={classes.root}>
        <div className={classes.content}>{children}</div>
      </Grid>
    </MainLayout>
  );
}
