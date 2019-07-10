import React, { PropsWithChildren } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Grid from "@material-ui/core/Grid";
import { LG_WIDTH, SM_WIDTH } from "../../constant/feed";

const useStyles = makeStyles(({ breakpoints }) => ({
  root: {
    display: "flex",
    flexDirection: "column",
    alignItems: "center"
  },
  content: {
    width: "100%",
    paddingLeft: 0,
    paddingRight: 0,
    [breakpoints.up("sm")]: {
      width: SM_WIDTH,
      paddingLeft: 12,
      paddingRight: 12
    },
    [breakpoints.up("md")]: {
      width: LG_WIDTH,
      paddingLeft: 12,
      paddingRight: 12
    },
    paddingTop: 70,
    paddingBottom: 20
  }
}));

interface Props {
  className?: string;
}

export default function Content({
  children,
  className
}: PropsWithChildren<Props>): JSX.Element {
  const classes = useStyles();
  return (
    <Grid item xs={12} className={classes.root}>
      <div className={`${classes.content} ${className ? className : ""}`}>
        {children}
      </div>
    </Grid>
  );
}
