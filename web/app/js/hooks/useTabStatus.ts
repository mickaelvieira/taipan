import { useEffect, useState } from "react";

export default function useTabStatus(): boolean {
  const [isVisible, setIsVisible] = useState(!document.hidden);

  useEffect(() => {
    function onVisibilityChange(): void {
      setIsVisible(!document.hidden);
    }

    document.addEventListener("visibilitychange", onVisibilityChange);

    return () => {
      document.removeEventListener("visibilitychange", onVisibilityChange);
    };
  }, []);

  return isVisible;
}
