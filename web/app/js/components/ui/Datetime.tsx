import React from "react";
import Typography from "@material-ui/core/Typography";

interface Props {
  label: string;
  value?: string;
}

export default React.memo(function Datetime({ label, value }: Props) {
  return (
    <Typography variant="body2">
      {label}: {value ? value : "Never"}
    </Typography>
  );
});
