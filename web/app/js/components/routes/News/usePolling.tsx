import { useEffect, useState } from "react";
import { useApolloClient } from "@apollo/react-hooks";
import useQueue, { Action } from "./useQueue";
import { FeedItem } from "../../../types/feed";
import query from "../../apollo/graphql/query/feeds/latest-news.graphql";

const WAITING_TIME = 30000;
const POLLING_FREQUENCY = 30000;
const POLLING_QUANTITY = 10;

export default function useLatest(firstId = ""): [FeedItem[], () => void] {
  const client = useApolloClient();
  const [shouldWait, setShouldWait] = useState(true);
  const [toId, setToId] = useState(firstId);
  const [state, dispatch] = useQueue();
  const { queue } = state;

  // console.log(firstId)
  // console.log(toId)
  // console.log(shouldWait)
  // console.log(queue)
  // console.log(queue.isFull())

  useEffect(() => {
    let timeout: number | undefined = undefined;
    let interval: number | undefined = undefined;

    function stopPolling(): void {
      if (interval) {
        window.clearInterval(interval);
      }
    }

    function stopWaiting(): void {
      if (timeout) {
        window.clearTimeout(timeout);
      }
    }

    async function poll(): Promise<null> {
      const result = await client.query({
        query,
        fetchPolicy: "no-cache",
        variables: {
          pagination: {
            limit: POLLING_QUANTITY,
            to: toId,
          },
        },
      });

      const { data } = result;
      const { results, last } = data.feeds.latestNews;
      if (results.length === 0) {
        return null;
      }

      dispatch([Action.PUSH, results]);
      // the results are in an ascending order
      // so the last ID is the greatest ID
      setToId(last);

      return null;
    }

    if (shouldWait) {
      // We need to wait before starting polling
      // otherwise it is overwhelming for the user
      timeout = window.setTimeout(() => setShouldWait(false), WAITING_TIME);
    } else if (queue.isFull()) {
      // No need to load the queue with hundreds of documents
      stopPolling();
    } else if (toId) {
      // we have the greatest ID we can start polling
      // documents having a greater ID
      interval = window.setInterval(poll, POLLING_FREQUENCY);
    } else {
      interval = window.setInterval(poll, POLLING_FREQUENCY);
      // we stop polling if we don't have the greatest ID
      // stopPolling();
    }

    return () => {
      stopPolling();
      stopWaiting();
    };
  }, [client, dispatch, queue, shouldWait, toId]);

  return [queue.items, () => dispatch([Action.RESET])];
}
