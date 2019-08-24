import { useReducer, Reducer, Dispatch } from "react";

interface State {
  terms: string[];
  hidden: boolean;
  paused: boolean;
}

type Payload = string[] | boolean;

export enum Action {
  TERMS = "terms",
  HIDDEN = "hidden",
  PAUSED = "paused"
}

function reducer(state: State, [type, payload]: [Action, Payload]): State {
  switch (type) {
    case Action.TERMS:
      return {
        ...state,
        terms: payload as string[]
      };
    case Action.HIDDEN:
      return {
        ...state,
        hidden: payload as boolean
      };
    case Action.PAUSED:
      return {
        ...state,
        paused: payload as boolean
      };
    default:
      throw new Error(`Invalid action type '${type}'`);
  }
}

type SearchReducer = Reducer<State, [Action, Payload]>;

export default function Search(): [State, Dispatch<[Action, Payload]>] {
  const [state, dispatch] = useReducer<SearchReducer>(reducer, {
    terms: [],
    hidden: false,
    paused: false
  });

  return [state, dispatch];
}
