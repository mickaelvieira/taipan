import React from "react";
import Typography from "@material-ui/core/Typography";

interface Props {
  value: string;
}

export default function PageTitle({ value }: Props): JSX.Element {
  return (
    <Typography component="h2" variant="h5">
      {value}
    </Typography>
  );
}
