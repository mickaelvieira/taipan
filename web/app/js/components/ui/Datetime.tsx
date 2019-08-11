import React from "react";
import dayjs from "dayjs";
import { Datetime as DatetimeType } from "../../types/scalars";

interface Props {
  value: DatetimeType;
  className?: string;
  format?: string;
}

export default React.memo(function Datetime({
  value,
  className,
  format = "DD MMMM YYYY, HH:mm:ss"
}: Props): JSX.Element {
  const date = dayjs(value);
  return (
    <span className={className ? className : ""}>{date.format(format)}</span>
  );
});
