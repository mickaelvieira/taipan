import { useEffect, useState } from "react";

export default function useWindowScroll(): boolean {
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
      setIsScrolling(true);
      timeout = window.setTimeout(onScrollStop, 200);
    }

    window.addEventListener("scroll", onScrollHandler);

    return () => {
      clearTimer();
      window.removeEventListener("scroll", onScrollHandler);
    };
  }, []);

  return isScrolling;
}
