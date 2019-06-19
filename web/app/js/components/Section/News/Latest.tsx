import React, { useEffect, useReducer, Reducer } from "react";
import { cloneDeep } from "lodash";
import { ApolloClient } from "apollo-client";
import { withApollo, WithApolloClient } from "react-apollo";
import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import { queryNews } from "../../apollo/Query/Feed";
import { query } from "../../apollo/Query/LatestNews";
import { getBoundaries } from "../../apollo/helpers/feed";
import { FeedItem } from "../../../types/feed";

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

function updateFeed(client: ApolloClient<object>, documents: FeedItem[]) {
  try {
    const data = client.readQuery({ query: queryNews });
    if (data) {
      const cloned = cloneDeep(data.News);
      const total = cloned.total + documents.length;
      const results = cloned.results;

      // Append documents at the top of the feed
      documents.reverse();
      results.unshift(...documents);
      const [first, last] = getBoundaries(results);

      client.writeQuery({
        query: queryNews,
        data: {
          News: {
            ...cloned,
            first,
            last,
            total,
            results
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
  documents: []
};

interface State {
  fromId: string;
  toId: string;
  documents: FeedItem[];
}

type Payload = string | FeedItem[] | undefined;

enum QueueActions {
  SET_FETCH_UNTIL = "to_id",
  SET_FETCH_FROM = "from_id",
  PUSH_DOCUMENTS = "documents",
  RESET = "reset"
}

function reducer(state: State, [type, payload]: [QueueActions, Payload]) {
  switch (type) {
    case QueueActions.SET_FETCH_UNTIL:
      return {
        ...state,
        toId: payload as string
      };
    case QueueActions.SET_FETCH_FROM:
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
      throw new Error();
  }
}

interface Props {
  firstId?: string;
  lastId?: string;
}

export default withApollo(function Latest({
  client,
  firstId = ""
}: WithApolloClient<Props>) {
  const classes = useStyles();
  const [state, dispatch] = useReducer<Reducer<State, [QueueActions, Payload]>>(
    reducer,
    initialState
  );
  const { toId, documents } = state;

  useEffect(() => {
    // We don't have yet the greatest ID in the queue
    // so we take the greatest ID in the feed
    // which is the first ID since the collection is in a descending order
    if (toId == "" && firstId !== "") {
      dispatch([QueueActions.SET_FETCH_UNTIL, firstId]);
    }
  }, [firstId, toId]);

  useEffect(() => {
    const FREQUENCY = 5000;
    let timeout: number | undefined = undefined;

    function clearTimer() {
      if (timeout) {
        window.clearTimeout(timeout);
      }
    }

    async function fetchData() {
      const LIMIT = 10;
      const result = await client.query({
        query,
        fetchPolicy: "no-cache",
        variables: {
          pagination: {
            limit: LIMIT,
            to: toId
          }
        }
      });

      const { data } = result;
      const { results, last } = data.LatestNews;
      if (results.length === 0) {
        return;
      }

      dispatch([QueueActions.PUSH_DOCUMENTS, results]);
      // the results are in an ascending order
      // so the last ID is the greatest ID
      dispatch([QueueActions.SET_FETCH_UNTIL, last]);
    }

    if (toId) {
      timeout = window.setInterval(fetchData, FREQUENCY);
    }

    return () => {
      clearTimer();
    };
  }, [client, toId]);

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
});
