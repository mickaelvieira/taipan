import { useEffect, useState } from "react";

export default function useScrollStatus(): boolean {
  const [isScrolling, setIsScrolling] = useState(false);

  useEffect(() => {
    let timeout: number | undefined = undefined;

    function clearTimer(): void {
      if (timeout) {
        window.clearTimeout(timeout);
      }
    }

    function onScrollStop(): void {
      setIsScrolling(false);
    }

    function onScrollHandler(): void {
      clearTimer();
      if (!isScrolling) {
        setIsScrolling(true);
      }
      timeout = window.setTimeout(onScrollStop, 200);
    }

    window.addEventListener("scroll", onScrollHandler);

    return () => {
      clearTimer();
      window.removeEventListener("scroll", onScrollHandler);
    };
  }, [isScrolling]);

  return isScrolling;
}
