import { useEffect, useState, RefObject } from "react";
import { isInViewport } from "../helpers/window";

export default function useIsInViewport(ref: RefObject<HTMLElement>): boolean {
  const [isVisible, setIsVisible] = useState(isInViewport(ref.current));

  useEffect(() => {
    let timeout: number | undefined = undefined;

    function clearTimer(): void {
      if (timeout) {
        window.clearTimeout(timeout);
      }
    }

    function onScrollStop(): void {
      if (!isVisible) {
        setIsVisible(isInViewport(ref.current));
      }
    }

    function onScrollHandler(): void {
      clearTimer();
      if (!isVisible) {
        setIsVisible(isInViewport(ref.current));
        timeout = window.setTimeout(onScrollStop, 400);
      } else {
        window.removeEventListener("scroll", onScrollHandler);
      }
    }

    window.addEventListener("scroll", onScrollHandler);

    setIsVisible(isInViewport(ref.current));

    return () => {
      clearTimer();
      window.removeEventListener("scroll", onScrollHandler);
    };
  }, [ref, isVisible]);

  return isVisible;
}
