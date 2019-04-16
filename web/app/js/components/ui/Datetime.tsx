import moment from "moment";
import React from "react";

interface Props {
  label: string;
  value?: string;
}

export default React.memo(function Datetime({ label, value }: Props) {
  return (
    <>
      <span className="datetime-label info-label">{label}</span>
      <span className="datetime-value info-value">
        {value ? moment(value).fromNow() : "Never"}
      </span>
    </>
  );
});
