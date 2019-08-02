import { useEffect, useState } from "react";
import { hasReachedTheBottom } from "../helpers/window";

export default function useWindowBottom(gap?: number): boolean {
  const [atTheBotttom, setIsAtTheBotttom] = useState(false);

  useEffect(() => {
    let timeout: number | undefined = undefined;

    function clearTimer(): void {
      if (timeout) {
        window.clearTimeout(timeout);
      }
    }

    function onScrollStop(): void {
      setIsAtTheBotttom(hasReachedTheBottom(gap));
    }

    function onScrollHandler(): void {
      clearTimer();
      setIsAtTheBotttom(hasReachedTheBottom(gap));
      timeout = window.setTimeout(onScrollStop, 400);
    }

    window.addEventListener("scroll", onScrollHandler);

    return () => {
      clearTimer();
      window.removeEventListener("scroll", onScrollHandler);
    };
  }, [gap]);

  return atTheBotttom;
}
