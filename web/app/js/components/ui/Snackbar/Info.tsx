import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import InfoIcon from "@material-ui/icons/InfoOutlined";
import Snackbar, { SnackbarProps } from "@material-ui/core/Snackbar";
import { MessageInfo } from "../../../types";
import Link from "@material-ui/core/Link";

const useStyles = makeStyles({
  icon: {
    marginRight: 12
  },
  message: {
    display: "flex",
    alignItems: "center"
  },
  link: {
    cursor: "pointer"
  }
});

interface Props extends SnackbarProps {
  info: MessageInfo | null;
  forceClose?: () => void;
}

export default function SnackbarInfo({
  info,
  forceClose,
  ...rest
}: Props): JSX.Element | null {
  const classes = useStyles();
  const action = (): void => {
    if (typeof forceClose === "function") {
      forceClose();
    }
    if (info && typeof info.action === "function") {
      info.action();
    }
  };

  return (
    <Snackbar
      anchorOrigin={{ vertical: "bottom", horizontal: "center" }}
      ContentProps={{
        "aria-describedby": "message-snackbar"
      }}
      message={
        <span id="message-snackbar" className={classes.message}>
          <InfoIcon className={classes.icon} />
          {info && (
            <>
              {info.message}
              {info.action && (
                <Link className={classes.link} onClick={action}>
                  {info.label ? info.label : "missing"}
                </Link>
              )}
            </>
          )}
        </span>
      }
      {...rest}
    />
  );
}
