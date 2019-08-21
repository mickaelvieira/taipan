import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import Snackbar, { SnackbarProps } from "@material-ui/core/Snackbar";
import SnackbarContent from "@material-ui/core/SnackbarContent";

const useStyles = makeStyles(({ palette, spacing, breakpoints }) => ({
  bar: {
    maxWidth: 600,
    backgroundColor: palette.primary.light,
    display: "flex",
    justifyContent: "center",
    margin: spacing(1),
    [breakpoints.up("md")]: {
      minWidth: 600
    }
  }
}));

export default function SnackbarWarning({
  className,
  message,
  action,
  ...rest
}: SnackbarProps): JSX.Element | null {
  const classes = useStyles();

  return (
    <Snackbar
      anchorOrigin={{ vertical: "top", horizontal: "center" }}
      {...rest}
    >
      <SnackbarContent
        className={`${classes.bar} ${className ? className : ""}`}
        message={<span>{message}</span>}
        action={action}
      />
    </Snackbar>
  );
}
