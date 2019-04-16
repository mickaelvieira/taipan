import React from "react";

interface Props {
  label?: string;
  value: URL;
}

export default React.memo(function Domain({ label = "Domain", value }: Props) {
  return (
    <>
      {/* <span className="domain-label info-label">{label}</span> */}
      <span className="domain-value info-value">
        <a href={`${value}`} title={`${value}`} target="_blank">
          {value.host}
        </a>
      </span>
    </>
  );
});
