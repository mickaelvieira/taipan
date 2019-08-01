import { useEffect, useState } from "react";

export default function useFeedScrolling(): number {
  const [height, setHeight] = useState(0);

  useEffect(() => {
    let timeout: number | undefined = undefined;

    function clearTimer(): void {
      if (timeout) {
        window.clearTimeout(timeout);
      }
    }

    function onScrollStop(): void {
      setHeight(document.documentElement.getBoundingClientRect().top);
    }

    function onScrollHandler(): void {
      clearTimer();
      timeout = window.setTimeout(onScrollStop, 400);
    }

    window.addEventListener("scroll", onScrollHandler);

    return () => {
      clearTimer();
      window.removeEventListener("scroll", onScrollHandler);
    };
  }, []);

  return height;
}
