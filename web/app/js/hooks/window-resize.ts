import { useEffect, useState } from "react";

export default function useWindowResize(): boolean {
  const [isResizing, setIsResizing] = useState(false);

  useEffect(() => {
    let timeout: number | undefined = undefined;

    function clearTimer(): void {
      if (timeout) {
        window.clearTimeout(timeout);
      }
    }

    function onResizeStop(): void {
      setIsResizing(false);
    }

    function onResizeHandler(): void {
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
