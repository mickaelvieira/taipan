import { useEffect, useState } from "react";

export default function useWindowScroll() {
  const [isScrolling, setIsScrolling] = useState(false);

  useEffect(() => {
    let timeout: number | undefined = undefined;

    function clearTimer() {
      if (timeout) {
        window.clearTimeout(timeout);
      }
    }

    function onScrollStop() {
      setIsScrolling(false);
    }

    function onScrollHandler() {
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
