import React from "react";
import Typography from "@material-ui/core/Typography";
import { makeStyles } from "@material-ui/core/styles";

const useStyles = makeStyles(() => ({
  title: {
    padding: "16px 16px 0 16px",
  },
}));

interface Props {
  value: string;
}

export default function PageTitle({ value }: Props): JSX.Element {
  const classes = useStyles();
  return (
    <Typography component="h2" variant="h5" className={classes.title}>
      {value}
    </Typography>
  );
}
