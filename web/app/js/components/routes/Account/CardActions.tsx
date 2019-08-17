import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import CardActions, { CardActionsProps } from "@material-ui/core/CardActions";

const useStyles = makeStyles(() => ({
  element: {
    padding: 16,
    justifyContent: "flex-end"
  }
}));

export default function ProfileCardActions({
  className,
  ...rest
}: CardActionsProps): JSX.Element | null {
  const classes = useStyles();
  return (
    <CardActions
      className={`${classes.element} ${className ? className : ""}`}
      {...rest}
    />
  );
}
