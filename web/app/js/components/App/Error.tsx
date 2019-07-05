import React, { ErrorInfo } from "react";

interface Props {
  error: Error | null;
  info: ErrorInfo | null;
}

export default function ErrorMessage({ error, info }: Props): JSX.Element {
  return (
    <div>
      <h2>Something went wrong.</h2>
      <details style={{ whiteSpace: "pre-wrap" }}>
        {error && error.toString()}
        <br />
        {info && info.componentStack}
      </details>
    </div>
  );
}
