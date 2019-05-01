import React, { PropsWithChildren } from "react";
import { withStyles, WithStyles, createStyles } from "@material-ui/core/styles";

const styles = () =>
  createStyles({
    scrolling: {
      pointerEvents: "none"
    },
    idle: {
      pointerEvents: "auto"
    }
  });

interface Props extends PropsWithChildren<WithStyles<typeof styles>> {
  isScrolling: boolean;
}

export default withStyles(styles)(function PointerEvents({
  children,
  isScrolling,
  classes
}: Props) {
  return (
    <div className={isScrolling ? classes.scrolling : classes.idle}>
      {children}
    </div>
  );
});
