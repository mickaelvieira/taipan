import React, { useEffect, PropsWithChildren } from "react";

export default function ScrollToTop({ children }: PropsWithChildren<{}>) {
  useEffect(() => {
    // Update the document title using the browser API
    // document.title = `You clicked ${count} times`;
    console.log("Scroll to TOP");
    window.scrollTo(0, 0);
  });

  return <>{children}</>;
}
