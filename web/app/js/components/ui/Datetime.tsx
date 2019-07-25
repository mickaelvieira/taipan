import React from "react";
import dayjs from "dayjs";
import { Datetime as DatetimeType } from "../../types/scalars";

interface Props {
  value: DatetimeType;
  className?: string;
}

export default React.memo(function Datetime({
  value,
  className
}: Props): JSX.Element {
  const date = dayjs(value);
  return (
    <span className={className ? className : ""}>
      {date.format("DD MMMM YYYY, HH:mm:ss")}
    </span>
  );
});
