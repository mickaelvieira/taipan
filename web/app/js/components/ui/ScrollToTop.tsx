import React, { useEffect, PropsWithChildren } from "react";

export default function ScrollToTop({
  children
}: PropsWithChildren<{}>): JSX.Element {
  useEffect(() => {
    window.scrollTo(0, 0);
  });

  return <>{children}</>;
}
