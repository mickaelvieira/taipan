import { useEffect, useState } from "react";
import { hasReachedTheBottom } from "../helpers/window";

export default function useWindowBottom(gap?: number): boolean {
  const [atTheBotttom, setIsAtTheBottom] = useState(false);

  useEffect(() => {
    let scrollStopTimeout: number | undefined = undefined;
    let scrollTimeout: number | undefined = undefined;

    function calculate(): void {
      setIsAtTheBottom(hasReachedTheBottom(gap));
    }

    function clearScrollStopTimer(): void {
      window.clearTimeout(scrollStopTimeout);
      scrollStopTimeout = undefined;
    }

    function clearScrollTimer(): void {
      window.clearTimeout(scrollTimeout);
      scrollTimeout = undefined;
    }

    function onScroll(): void {
      clearScrollTimer();
      calculate();
    }

    function onScrollStop(): void {
      clearScrollTimer();
      calculate();
    }

    function startScrollTimer(): void {
      if (!scrollTimeout) {
        scrollTimeout = window.setTimeout(onScroll, 400);
      }
    }

    function startScrollStopTimer(): void {
      if (!scrollStopTimeout) {
        scrollStopTimeout = window.setTimeout(onScrollStop, 200);
      }
    }

    function onScrollHandler(): void {
      clearScrollStopTimer();
      startScrollTimer();
      startScrollStopTimer();
    }

    window.addEventListener("scroll", onScrollHandler);

    calculate();

    return () => {
      clearScrollTimer();
      clearScrollStopTimer();
      window.removeEventListener("scroll", onScrollHandler);
    };
  }, [gap]);

  return atTheBotttom;
}
