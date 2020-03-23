import { useReducer, Reducer, Dispatch } from "react";

interface State {
  terms: string[];
  showDeleted: boolean;
  pausedOnly: boolean;
}

type Payload = string[] | boolean;

export enum Action {
  TERMS = "terms",
  DELETED = "deleted",
  PAUSED = "paused",
}

function reducer(state: State, [type, payload]: [Action, Payload]): State {
  switch (type) {
    case Action.TERMS:
      return {
        ...state,
        terms: payload as string[],
      };
    case Action.DELETED:
      return {
        ...state,
        showDeleted: payload as boolean,
      };
    case Action.PAUSED:
      return {
        ...state,
        pausedOnly: payload as boolean,
      };
    default:
      throw new Error(`Invalid action type '${type}'`);
  }
}

type SearchReducer = Reducer<State, [Action, Payload]>;

export default function Search(): [State, Dispatch<[Action, Payload]>] {
  const [state, dispatch] = useReducer<SearchReducer>(reducer, {
    terms: [],
    showDeleted: false,
    pausedOnly: false,
  });

  return [state, dispatch];
}
