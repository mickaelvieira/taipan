import { useEffect, useState, RefObject } from "react";
import { willBeSoonInViewport } from "../helpers/window";

export default function useWillBeSoonInViewport(ref: RefObject<HTMLElement>) {
  const [isInArea, setIsInArea] = useState(willBeSoonInViewport(ref.current));

  useEffect(() => {
    let timeout: number | undefined = undefined;

    function clearTimer() {
      if (timeout) {
        window.clearTimeout(timeout);
      }
    }

    function onScrollStop() {
      if (!isInArea) {
        setIsInArea(willBeSoonInViewport(ref.current));
      }
    }

    function onScrollHandler() {
      clearTimer();
      if (!isInArea) {
        setIsInArea(willBeSoonInViewport(ref.current));
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
  }, [ref, isInArea]);

  return isInArea;
}
