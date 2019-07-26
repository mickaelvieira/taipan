import React from "react";
import IconButton, { IconButtonProps } from "@material-ui/core/IconButton";
import CircularProgress from "@material-ui/core/CircularProgress";
import { SvgIconProps } from "@material-ui/core/SvgIcon";

export interface ButtonBaseProps extends IconButtonProps {
  Icon: React.ComponentType<SvgIconProps>;
  label: string;
  textOnly?: boolean;
  iconOnly?: boolean;
  isLoading?: boolean;
}

export default function ButtonBase({
  Icon,
  label,
  isLoading = false,
  textOnly = false,
  iconOnly = false,
  ...rest
}: ButtonBaseProps): JSX.Element {
  return (
    <IconButton {...rest}>
      {!textOnly && (isLoading ? <CircularProgress size={16} /> : <Icon />)}
      {!iconOnly && <span>{label}</span>}
    </IconButton>
  );
}
