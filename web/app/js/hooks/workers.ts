// import { useEffect } from "react";
// import workers from "../services/workers";
// import { Bookmark } from "types/bookmark";

// export default function useWorkers(onReceived: (data: Bookmark[]) => void) {
//   function onMessageReceived(event: MessageEvent) {
//     onReceived(event.data);
//   }

//   useEffect(() => {
//     workers.fetchWorker.addEventListener("message", onMessageReceived);
//     return () => {
//       workers.fetchWorker.removeEventListener("message", onMessageReceived);
//     };
//   });

//   return workers.fetchWorker;
// }
