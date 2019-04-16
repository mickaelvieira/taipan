import React from "react";

interface Props {
  label: string;
  value: boolean;
}

export default React.memo(function YesNo({ label, value }: Props) {
  return (
    <>
      <span className="domain-label info-label">{label}</span>
      <span className="domain-value info-value">{value ? "Yes" : "No"}</span>
    </>
  );
});
