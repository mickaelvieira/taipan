import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import MuiInputBase, { InputBaseProps } from "@material-ui/core/InputBase";

const useStyles = makeStyles(({ spacing, palette }) => ({
  element: {
    padding: `${spacing(0.5)}px ${spacing(1)}px`,
    border: `1px solid ${palette.grey[400]}`,
    borderRadius: 4,
    ["&:focus-within"]: {
      borderColor: palette.primary.main
    }
  },
  input: {
    ["&:focus"]: {
      color: palette.primary.main
    }
  }
}));

export default function InputBase({
  className,
  ...rest
}: InputBaseProps): JSX.Element {
  const classes = useStyles();
  return (
    <MuiInputBase
      className={`${classes.element} ${className ? className : ""}`}
      inputProps={{
        className: classes.input
      }}
      {...rest}
    />
  );
}
