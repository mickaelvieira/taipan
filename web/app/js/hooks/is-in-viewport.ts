import { useEffect, useState, RefObject } from "react";
import { isInViewport } from "../helpers/window";

export default function useIsVisible(ref: RefObject<HTMLElement>) {
  const [isVisible, setIsVisible] = useState(isInViewport(ref.current));

  useEffect(() => {
    let timeout: number | undefined = undefined;

    function clearTimer() {
      if (timeout) {
        window.clearTimeout(timeout);
      }
    }

    function onScrollStop() {
      if (!isVisible) {
        setIsVisible(isInViewport(ref.current));
      }
    }

    function onScrollHandler() {
      clearTimer();
      if (!isVisible) {
        setIsVisible(isInViewport(ref.current));
        timeout = window.setTimeout(onScrollStop, 400);
      } else {
        window.removeEventListener("scroll", onScrollHandler);
      }
    }

    window.addEventListener("scroll", onScrollHandler);

    return () => {
      clearTimer();
      window.removeEventListener("scroll", onScrollHandler);
    };
  }, [ref, isVisible]);

  return isVisible;
}
