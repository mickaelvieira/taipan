import React from "react";
import {
  withStyles,
  WithStyles,
  createStyles,
  Theme
} from "@material-ui/core/styles";
import CircularProgress from "@material-ui/core/CircularProgress";

const styles = ({ spacing }: Theme) =>
  createStyles({
    container: {
      flex: 1,
      display: "flex",
      justifyContent: "center",
      alignItems: "center"
    },
    progress: {
      margin: spacing.unit * 2
    }
  });

export default withStyles(styles)(function Loader({
  classes
}: WithStyles<typeof styles>) {
  return (
    <div className={classes.container}>
      <CircularProgress className={classes.progress} />
    </div>
  );
});
