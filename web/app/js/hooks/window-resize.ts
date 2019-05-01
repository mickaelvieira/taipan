import { useEffect, useState } from "react";

export default function useWindowResize() {
  const [isResizing, setIsResizing] = useState(false);

  useEffect(() => {
    let timeout: number | undefined = undefined;

    function clearTimer() {
      if (timeout) {
        window.clearTimeout(timeout);
      }
    }

    function onResizeStop() {
      setIsResizing(false);
    }

    function onResizeHandler() {
      clearTimer();
      setIsResizing(true);
      timeout = window.setTimeout(onResizeStop, 200);
    }

    window.addEventListener("resize", onResizeHandler);

    return () => {
      clearTimer();
      window.removeEventListener("resize", onResizeHandler);
    };
  }, []);

  return isResizing;
}
