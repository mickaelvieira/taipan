import { useEffect, useState } from "react";

export default function useTabStatus() {
  const [isVisible, setIsVisible] = useState(!document.hidden);

  function onVisibilityChange() {
    setIsVisible(!document.hidden);
  }

  useEffect(() => {
    document.addEventListener("visibilitychange", onVisibilityChange);
    return () => {
      document.removeEventListener("visibilitychange", onVisibilityChange);
    };
  }, []);

  return isVisible;
}
