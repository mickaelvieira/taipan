import React, { PropsWithChildren } from "react";
import { withStyles, WithStyles, createStyles } from "@material-ui/core/styles";

const styles = () =>
  createStyles({
    container: {
      marginBottom: 60,
      width: "100%",
      minHeight: "100vh"
    }
  });

export default withStyles(styles)(function FeedContainer({
  children,
  classes
}: PropsWithChildren<WithStyles<typeof styles>>) {
  return <section className={classes.container}>{children}</section>;
});
