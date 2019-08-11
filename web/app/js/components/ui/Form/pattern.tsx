import React from "react";
import { makeStyles } from "@material-ui/core/styles";

import FormLabel, { FormLabelProps } from "@material-ui/core/FormLabel";
const useStyles = makeStyles(({ palette }) => ({
  element: {
    border: `1px solid ${palette.grey[400]}`
  }
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
