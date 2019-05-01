import { useEffect, useState } from "react";
import { hasReachedTheBottom } from "../helpers/window";

export default function useWindoBottom() {
  const [atTheBotttom, setIsAtTheBotttom] = useState(false);

  useEffect(() => {
    let timeout: number | undefined = undefined;

    function clearTimer() {
      if (timeout) {
        window.clearTimeout(timeout);
      }
    }

    function onScrollStop() {
      setIsAtTheBotttom(hasReachedTheBottom());
    }

    function onScrollHandler() {
      clearTimer();
      setIsAtTheBotttom(hasReachedTheBottom());
      timeout = window.setTimeout(onScrollStop, 400);
    }

    window.addEventListener("scroll", onScrollHandler);

    return () => {
      clearTimer();
      window.removeEventListener("scroll", onScrollHandler);
    };
  }, []);

  return atTheBotttom;
}
