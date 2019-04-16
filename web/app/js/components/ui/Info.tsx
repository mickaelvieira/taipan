import React from "react";

interface Props {
  label: string;
  value: string;
}

export default React.memo(function Info({ label, value }: Props) {
  return (
    <>
      <span className="info-label">{label}</span>
      <span className="info-value">{value}</span>
    </>
  );
});
