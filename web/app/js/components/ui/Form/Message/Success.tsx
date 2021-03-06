import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import Typography, { TypographyProps } from "@material-ui/core/Typography";

const useStyles = makeStyles(({ spacing }) => ({
  element: {
    marginTop: spacing(2),
  },
}));

export default function Hint({
  className,
  ...rest
}: TypographyProps): JSX.Element {
  const classes = useStyles();
  return (
    <Typography
      component="p"
      className={`${classes.element} ${className ? className : ""}`}
      {...rest}
    />
  );
}
