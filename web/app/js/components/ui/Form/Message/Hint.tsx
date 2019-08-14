import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import FormHelperText, {
  FormHelperTextProps
} from "@material-ui/core/FormHelperText";

const useStyles = makeStyles(() => ({
  element: {}
}));

export default function Hint({
  className,
  ...rest
}: FormHelperTextProps): JSX.Element {
  const classes = useStyles();
  return (
    <FormHelperText
      className={`${classes.element} ${className ? className : ""}`}
      {...rest}
    />
  );
}
