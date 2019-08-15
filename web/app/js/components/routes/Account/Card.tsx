import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import Card, { CardProps } from "@material-ui/core/Card";

const useStyles = makeStyles(() => ({
  card: {
    marginTop: 16
  }
}));

export default function ProfileCard({
  className,
  ...rest
}: CardProps): JSX.Element {
  const classes = useStyles();
  return (
    <Card
      className={`${classes.card} ${className ? className : ""}`}
      {...rest}
    />
  );
}
