import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import FormLabel, { FormLabelProps } from "@material-ui/core/FormLabel";

const useStyles = makeStyles(({ spacing }) => ({
  element: {
    margin: `${spacing(1)}px 0`,
  },
}));

export default function Label({
  className,
  ...rest
}: FormLabelProps): JSX.Element {
  const classes = useStyles();
  return (
    <FormLabel
      className={`${classes.element} ${className ? className : ""}`}
      {...rest}
    />
  );
}
