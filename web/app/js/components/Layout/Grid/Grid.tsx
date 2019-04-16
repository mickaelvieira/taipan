import React from "react";

interface Props {
  children: any;
}

export default function Grid({ children }: Props) {
  return (
    <div className="app-grid">
      <div className="col" />
      <div className="col container">{children}</div>
      <div className="col" />
    </div>
  );
}
