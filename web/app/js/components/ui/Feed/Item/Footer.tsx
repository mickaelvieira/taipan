import React, { PropsWithChildren } from "react";
import { withStyles, WithStyles, createStyles } from "@material-ui/core/styles";

const styles = () =>
  createStyles({
    footer: {
      display: "flex",
      flexDirection: "row",
      justifyContent: "space-between"
    }
  });

export default withStyles(styles)(function ItemFooter({
  children,
  classes
}: PropsWithChildren<WithStyles<typeof styles>>) {
  return <footer className={classes.footer}>{children}</footer>;
});
