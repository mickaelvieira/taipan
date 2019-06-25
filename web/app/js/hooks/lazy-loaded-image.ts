import { useEffect, useState, RefObject } from "react";

const shouldBeShown = (element: HTMLElement | null): boolean => {
  if (!element) {
    return false;
  }

  const bounding = element.getBoundingClientRect();
  const bottom = window.innerHeight || document.documentElement.clientHeight;

  return bounding.top <= bottom + 400;
};

export default function useLazyLoadedImage(
  ref: RefObject<HTMLElement>
): boolean {
  const [isVisible, setIsVisible] = useState(shouldBeShown(ref.current));

  useEffect(() => {
    let timeout: number | undefined = undefined;

    function clearTimer(): void {
      if (timeout) {
        window.clearTimeout(timeout);
      }
    }

    function onScrollStop(): void {
      if (!isVisible) {
        setIsVisible(shouldBeShown(ref.current));
      }
    }

    function onScrollHandler(): void {
      clearTimer();
      if (!isVisible) {
        setIsVisible(shouldBeShown(ref.current));
        timeout = window.setTimeout(onScrollStop, 400);
      } else {
        window.removeEventListener("scroll", onScrollHandler);
      }
    }

    window.addEventListener("scroll", onScrollHandler);

    setIsVisible(shouldBeShown(ref.current));

    return () => {
      clearTimer();
      window.removeEventListener("scroll", onScrollHandler);
    };
  }, [ref, isVisible]);

  return isVisible;
}
