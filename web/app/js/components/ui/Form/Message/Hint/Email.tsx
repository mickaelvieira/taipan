import React from "react";
import BaseHint from "./Base";
import { FormHelperTextProps } from "@material-ui/core/FormHelperText";

export default function EmailHint(props: FormHelperTextProps): JSX.Element {
  return (
    <BaseHint {...props}>
      Make sure it&apos;s correct, we will send you a comfirmation email.
    </BaseHint>
  );
}
