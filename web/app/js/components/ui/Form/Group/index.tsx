import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import FormControl, { FormControlProps } from "@material-ui/core/FormControl";

const useStyles = makeStyles(({ spacing }) => ({
  element: {
    margin: `${spacing(1)}px 0 0`,
    display: "flex",
    flexDirection: "column",
  },
}));

export default function Group({
  className,
  ...rest
}: FormControlProps): JSX.Element {
  const classes = useStyles();
  return (
    <FormControl
      className={`${classes.element} ${className ? className : ""}`}
      {...rest}
    />
  );
}
