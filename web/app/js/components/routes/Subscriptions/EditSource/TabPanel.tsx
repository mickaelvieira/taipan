import React, { PropsWithChildren } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Box from "@material-ui/core/Box";
import Typography from "@material-ui/core/Typography";

const useStyles = makeStyles(() => ({
  paper: {
    display: "flex",
    flexDirection: "column"
  }
}));

interface Props {
  index: number;
  value: number;
  className?: string;
}

export default function TabPanel({
  children,
  index,
  value,
  className
}: PropsWithChildren<Props>): JSX.Element | null {
  const classes = useStyles();
  const styles = [classes.paper];

  if (className) {
    styles.push(className);
  }

  return (
    <Typography
      component="div"
      hidden={value !== index}
      role="tabpanel"
      id={`source-info-panel-${index}`}
      aria-labelledby={`source-info-tab-${index}`}
    >
      <Box className={`${styles.join(" ")}`}>{children}</Box>
    </Typography>
  );
}
