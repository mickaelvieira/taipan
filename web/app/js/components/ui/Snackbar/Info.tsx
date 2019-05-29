import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import InfoIcon from "@material-ui/icons/InfoOutlined";
import Snackbar, { SnackbarProps } from "@material-ui/core/Snackbar";

const useStyles = makeStyles({
  icon: {
    marginRight: 12
  },
  message: {
    display: "flex",
    alignItems: "center"
  }
});

interface Props extends SnackbarProps {
  info: string;
}

export default function SnackbarInfo({ info, ...rest }: Props) {
  const classes = useStyles();
  return (
    <Snackbar
      anchorOrigin={{ vertical: "bottom", horizontal: "center" }}
      ContentProps={{
        "aria-describedby": "offline-message-snackbar"
      }}
      message={
        <span id="offline-message-snackbar" className={classes.message}>
          <InfoIcon className={classes.icon} />
          {info}
        </span>
      }
      {...rest}
    />
  );
}
