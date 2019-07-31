export default function(window: Window): void {
  const url = new URL(`${window.location}`);
  if (url.protocol === "https:") {
    if ("serviceWorker" in window.navigator) {
      console.log("Trying to register the service worker...");
      window.navigator.serviceWorker
        .register("/sw.js", { scope: "/" })
        .then(registration => {
          console.log(`Registration succeeded. Scope is ${registration.scope}`);
        })
        .catch(error => {
          console.log(`Registration failed with ${error}`);
        });
    } else {
      console.warn("Service workers are not supported");
    }
  }
}
