import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import Typography from "@material-ui/core/Typography";

const useStyles = makeStyles(() => ({
  title: {
    margin: "24px 0",
  },
}));

interface Props {
  value: string;
}

export default function PageTitle({ value }: Props): JSX.Element {
  const classes = useStyles();

  return (
    <Typography component="h1" variant="h4" className={classes.title}>
      {value}
    </Typography>
  );
}
