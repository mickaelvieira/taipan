import React from "react";
import TabBase from "@material-ui/core/Tab";

interface Props {
  value: number;
  label: string;
}

export default function Tab({ value, label, ...rest }: Props): JSX.Element {
  return (
    <TabBase
      label={label}
      value={value}
      id={`source-info-tab-${value}`}
      aria-controls={`source-info-panel-${value}`}
      {...rest}
    />
  );
}
