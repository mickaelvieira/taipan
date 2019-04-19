import React from "react";
import { withStyles, WithStyles, createStyles } from "@material-ui/core/styles";
import InfoIcon from "@material-ui/icons/InfoOutlined";
import Snackbar from "@material-ui/core/Snackbar";

const styles = () =>
  createStyles({
    icon: {
      marginRight: 12
    },
    message: {
      display: "flex",
      alignItems: "center"
    }
  });

interface Props extends WithStyles<typeof styles> {
  isOpen: boolean;
  message: string;
}

export default withStyles(styles)(function SnackbarInfo({
  message,
  isOpen,
  classes
}: Props) {
  return (
    <Snackbar
      anchorOrigin={{ vertical: "bottom", horizontal: "center" }}
      open={isOpen}
      ContentProps={{
        "aria-describedby": "offline-message-snackbar"
      }}
      message={
        <span id="offline-message-snackbar" className={classes.message}>
          <InfoIcon className={classes.icon} />
          {message}
        </span>
      }
    />
  );
});
