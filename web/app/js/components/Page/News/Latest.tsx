import React, { useEffect, useReducer, Reducer } from "react";
import { cloneDeep } from "lodash";
import { ApolloClient } from "apollo-client";
import { useApolloClient } from "@apollo/react-hooks";
import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import query from "../../apollo/graphql/query/feeds/latest-news.graphql";
import { queryNews } from "../../apollo/Query/Feed";
import { getBoundaries } from "../../apollo/helpers/feed";
import { FeedItem, FeedQueryData } from "../../../types/feed";

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

const initialState = {
  fromId: "",
  toId: "",
  documents: [],
  shouldWait: true
};

interface State {
  fromId: string;
  toId: string;
  documents: FeedItem[];
  shouldWait: boolean;
}

type Payload = string | boolean | FeedItem[] | undefined;

enum QueueActions {
  SET_FETCH_TO_ID = "to_id",
  SET_FETCH_FROM_ID = "from_id",
  SET_SHOULD_WAIT = "waiting",
  PUSH_DOCUMENTS = "documents",
  RESET = "reset"
}

function reducer(
  state: State,
  [type, payload]: [QueueActions, Payload]
): State {
  switch (type) {
    case QueueActions.SET_SHOULD_WAIT:
      return {
        ...state,
        shouldWait: payload as boolean
      };
    case QueueActions.SET_FETCH_TO_ID:
      return {
        ...state,
        toId: payload as string
      };
    case QueueActions.SET_FETCH_FROM_ID:
      return {
        ...state,
        fromId: payload as string
      };
    case QueueActions.PUSH_DOCUMENTS:
      return {
        ...state,
        documents: [...state.documents, ...(payload as FeedItem[])]
      };
    case QueueActions.RESET:
      return { ...initialState };
    default:
      throw new Error(`Invalid action type '${type}'`);
  }
}

interface Props {
  firstId?: string;
  lastId?: string;
}

export default function Latest({
  firstId = ""
}: Props) {
  const classes = useStyles();
  const client = useApolloClient()
  const [state, dispatch] = useReducer<Reducer<State, [QueueActions, Payload]>>(
    reducer,
    initialState
  );
  const { shouldWait, toId, documents } = state;

  useEffect(() => {
    // We don't have yet the greatest ID in the queue
    // so we take the greatest ID in the feed
    // which is the first ID since the collection is in a descending order
    if (toId == "" && firstId !== "") {
      dispatch([QueueActions.SET_FETCH_TO_ID, firstId]);
    }
  }, [firstId, toId]);

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

      dispatch([QueueActions.PUSH_DOCUMENTS, results]);
      // the results are in an ascending order
      // so the last ID is the greatest ID
      dispatch([QueueActions.SET_FETCH_TO_ID, last]);

      return null;
    }

    if (shouldWait) {
      // We need to wait before starting polling
      // otherwise it is overwhelming for the user
      timeout = window.setTimeout(
        () => dispatch([QueueActions.SET_SHOULD_WAIT, false]),
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
  }, [client, documents.length, shouldWait, toId]);

  return documents.length === 0 ? null : (
    <div className={classes.container}>
      <Button
        className={classes.button}
        onClick={() => {
          // append documents to the feed
          updateFeed(client, documents);
          // reset the queue
          dispatch([QueueActions.RESET, undefined]);
          window.scroll(0, 0);
        }}
      >
        See {documents.length} latest news
      </Button>
    </div>
  );
}
