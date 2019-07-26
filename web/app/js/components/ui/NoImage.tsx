import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import IconImage from "@material-ui/icons/ImageOutlined";

const useStyles = makeStyles(({ palette }) => ({
  container: {
    display: "flex",
    flexDirection: "column",
    justifyContent: "center",
    alignItems: "center",
    minHeight: 300,
    backgroundColor: palette.grey[200]
  },
  icon: {
    fontSize: "10rem",
    color: palette.grey[400]
  }
}));

interface Props {
  className?: string;
}

export default React.memo(function NoImage({ className }: Props) {
  const classes = useStyles();
  return (
    <div
      className={`${classes.container} ${
        typeof className === "string" ? className : ""
      }`}
    >
      <IconImage className={classes.icon} />
    </div>
  );
});
