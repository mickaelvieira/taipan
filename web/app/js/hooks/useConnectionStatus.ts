import { useEffect, useState } from "react";

export default function useConnectionStatus(): boolean {
  const [isOnline, setIsOnline] = useState(window.navigator.onLine);

  useEffect(() => {
    function onStatusChange(): void {
      setIsOnline(window.navigator.onLine);
    }

    window.addEventListener("online", onStatusChange);
    window.addEventListener("offline", onStatusChange);

    return () => {
      window.removeEventListener("online", onStatusChange);
      window.removeEventListener("offline", onStatusChange);
    };
  }, []);

  return isOnline;
}
