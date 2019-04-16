export default function(navigator) {
  if ("serviceWorker" in navigator) {
    console.log("Trying to register the service worker...");
    navigator.serviceWorker
      .register("/sw.js", { useCache: false })
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
