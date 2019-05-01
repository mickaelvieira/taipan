import { useEffect, useState } from "react";

export default function useTabStatus() {
  const [isVisible, setIsVisible] = useState(!document.hidden);

  useEffect(() => {
    function onVisibilityChange() {
      setIsVisible(!document.hidden);
    }

    document.addEventListener("visibilitychange", onVisibilityChange);

    return () => {
      document.removeEventListener("visibilitychange", onVisibilityChange);
    };
  }, []);

  return isVisible;
}
