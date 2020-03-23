import React, { useState } from "react";
import { makeStyles } from "@material-ui/core/styles";
import IconButton from "@material-ui/core/IconButton";
import VisibleIcon from "@material-ui/icons/Visibility";
import HiddenIcon from "@material-ui/icons/VisibilityOff";
import { InputBaseProps } from "@material-ui/core/InputBase";
import InputBase from "./Base";

const useStyles = makeStyles(({ spacing, palette }) => ({
  wrapper: {
    display: "flex",
    ["&:focus-within .revealer"]: {
      borderColor: palette.primary.main,
    },
  },
  element: {
    flexGrow: 1,
    borderRight: 0,
    borderTopRightRadius: 0,
    borderBottomRightRadius: 0,
  },
  button: {
    padding: spacing(1),
    border: `1px solid ${palette.grey[400]}`,
    borderLeft: 0,
    borderRadius: 4,
    borderTopLeftRadius: 0,
    borderBottomLeftRadius: 0,
  },
}));

export default function InputPassword({
  className,
  ...rest
}: InputBaseProps): JSX.Element {
  const classes = useStyles();
  const [isVisible, setVisibility] = useState(false);

  return (
    <div className={classes.wrapper}>
      <InputBase
        type={isVisible ? "text" : "password"}
        autoComplete="on"
        autoCapitalize="off"
        autoCorrect="off"
        className={`${classes.element} ${className ? className : ""}`}
        {...rest}
      />
      <IconButton
        onClick={() => setVisibility(!isVisible)}
        className={`revealer ${classes.button}`}
      >
        {isVisible ? <HiddenIcon /> : <VisibleIcon />}
      </IconButton>
    </div>
  );
}
