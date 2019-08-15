import React from "react";
import BaseHint from "./Base";
import { FormHelperTextProps } from "@material-ui/core/FormHelperText";

export default function PasswordHint(props: FormHelperTextProps): JSX.Element {
  return (
    <BaseHint {...props}>Make sure it&apos;s at least 10 characters.</BaseHint>
  );
}
