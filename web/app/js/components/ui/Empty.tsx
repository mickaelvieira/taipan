import React from "react";

interface Props {
  width?: number;
  height?: number;
}

export default function Empty({ width = 1, height = 1 }: Props): JSX.Element {
  return (
    <svg height={height} width={width} xmlns="http://www.w3.org/2000/svg">
      <rect fill="none" height={height} ry="0" width={width} x="0" y="0" />
    </svg>
  );
}
