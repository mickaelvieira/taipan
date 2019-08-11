import { useReducer, Reducer, Dispatch } from "react";
import { FeedItem } from "../../../types/feed";

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

export enum Action {
  SET_FETCH_TO_ID = "to_id",
  SET_FETCH_FROM_ID = "from_id",
  SET_SHOULD_WAIT = "waiting",
  PUSH_DOCUMENTS = "documents",
  RESET = "reset"
}

function reducer(state: State, [type, payload]: [Action, Payload]): State {
  switch (type) {
    case Action.SET_SHOULD_WAIT:
      return {
        ...state,
        shouldWait: payload as boolean
      };
    case Action.SET_FETCH_TO_ID:
      return {
        ...state,
        toId: payload as string
      };
    case Action.SET_FETCH_FROM_ID:
      return {
        ...state,
        fromId: payload as string
      };
    case Action.PUSH_DOCUMENTS:
      return {
        ...state,
        documents: [...state.documents, ...(payload as FeedItem[])]
      };
    case Action.RESET:
      return { ...initialState };
    default:
      throw new Error(`Invalid action type '${type}'`);
  }
}

export default function useLatestReducer(): [
  State,
  Dispatch<[Action, Payload]>
] {
  const [state, dispatch] = useReducer<Reducer<State, [Action, Payload]>>(
    reducer,
    initialState
  );

  return [state, dispatch];
}
