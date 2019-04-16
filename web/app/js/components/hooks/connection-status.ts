import { useEffect, useState } from "react";

export default function useConnectionStatus() {
  const [isOnline, setIsOnline] = useState(window.navigator.onLine);

  function onStatusChange() {
    setIsOnline(window.navigator.onLine);
  }

  useEffect(() => {
    window.addEventListener("online", onStatusChange);
    window.addEventListener("offline", onStatusChange);
    return () => {
      window.removeEventListener("online", onStatusChange);
      window.removeEventListener("offline", onStatusChange);
    };
  }, []);

  return isOnline;
}
