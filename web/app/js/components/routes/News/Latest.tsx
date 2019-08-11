import React, { useEffect } from "react";
import { cloneDeep } from "lodash";
import { ApolloClient } from "apollo-client";
import { useApolloClient } from "@apollo/react-hooks";
import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import query from "../../apollo/graphql/query/feeds/latest-news.graphql";
import { queryNews } from "../../apollo/Query/Feed";
import { getBoundaries } from "../../apollo/helpers/feed";
import { FeedItem, FeedQueryData } from "../../../types/feed";
import useLatestReducer, { Action } from "./useLatestReducer";

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    container: {
      position: "fixed",
      backgroundColor: "#fff",
      alignSelf: "center"
    },
    button: {
      margin: theme.spacing(1)
    }
  })
);

// @TODO move this to feed updater
function updateFeed(client: ApolloClient<object>, documents: FeedItem[]): void {
  try {
    const data = client.readQuery({ query: queryNews }) as FeedQueryData;
    if (data) {
      const cloned = cloneDeep(data.feeds.news);
      const total = cloned.total + documents.length;
      const results = cloned.results;

      // Append documents at the top of the feed
      documents.reverse();
      results.unshift(...documents);
      const [first, last] = getBoundaries(results);

      client.writeQuery({
        query: queryNews,
        data: {
          feeds: {
            ...data.feeds,
            news: {
              ...cloned,
              first,
              last,
              total,
              results
            }
          }
        }
      });
    }
  } catch (e) {
    console.warn(e);
  }
}

interface Props {
  firstId?: string;
  lastId?: string;
}

export default function Latest({ firstId = "" }: Props): JSX.Element | null {
  const classes = useStyles();
  const client = useApolloClient();
  const [state, dispatch] = useLatestReducer();
  const { shouldWait, toId, documents } = state;

  useEffect(() => {
    // We don't have yet the greatest ID in the queue
    // so we take the greatest ID in the feed
    // which is the first ID since the collection is in a descending order
    if (toId == "" && firstId !== "") {
      dispatch([Action.SET_FETCH_TO_ID, firstId]);
    }
  }, [dispatch, firstId, toId]);

  useEffect(() => {
    const WAITING_TIME = 30000;
    const POLLING_FREQUENCY = 10000;
    const POLLING_QUANTITY = 10;
    const MAX_QUEUE_LENGTH = 50;
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
            to: toId
          }
        }
      });

      const { data } = result;
      const { results, last } = data.feeds.latestNews;
      if (results.length === 0) {
        return null;
      }

      dispatch([Action.PUSH_DOCUMENTS, results]);
      // the results are in an ascending order
      // so the last ID is the greatest ID
      dispatch([Action.SET_FETCH_TO_ID, last]);

      return null;
    }

    if (shouldWait) {
      // We need to wait before starting polling
      // otherwise it is overwhelming for the user
      timeout = window.setTimeout(
        () => dispatch([Action.SET_SHOULD_WAIT, false]),
        WAITING_TIME
      );
    } else if (documents.length >= MAX_QUEUE_LENGTH) {
      // No need to load the queue with hundreds of documents
      stopPolling();
    } else if (toId) {
      // we have the greatest ID we can start polling
      // documents having a greater ID
      interval = window.setInterval(poll, POLLING_FREQUENCY);
    } else {
      // we stop polling if we don't have the greatest ID
      stopPolling();
    }

    return () => {
      stopPolling();
      stopWaiting();
    };
  }, [client, dispatch, documents.length, shouldWait, toId]);

  return documents.length === 0 ? null : (
    <div className={classes.container}>
      <Button
        className={classes.button}
        onClick={() => {
          // append documents to the feed
          updateFeed(client, documents);
          // reset the queue
          dispatch([Action.RESET, undefined]);
          window.scroll(0, 0);
        }}
      >
        See {documents.length} latest news
      </Button>
    </div>
  );
}
