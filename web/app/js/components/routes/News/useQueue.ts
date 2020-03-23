import { useReducer, Reducer, Dispatch } from "react";
import { FeedItem } from "../../../types/feed";

const MAX_QUEUE_LENGTH = 50;

class Queue {
  items: FeedItem[] = [];

  constructor(items: FeedItem[] = []) {
    this.items = items;
  }

  push(items: FeedItem[]): Queue {
    return new Queue([...this.items, ...items]);
  }

  isFull(): boolean {
    return this.items.length >= MAX_QUEUE_LENGTH;
  }
}

interface State {
  queue: Queue;
}

type Payload = FeedItem[];

type DispatchArgs = [Action, Payload?];

export enum Action {
  PUSH = "documents",
  RESET = "reset",
}

const initialState = {
  queue: new Queue(),
};

function reducer(state: State, [type, payload]: DispatchArgs): State {
  switch (type) {
    case Action.PUSH:
      return {
        ...state,
        queue: state.queue.push(payload as FeedItem[]),
      };
    case Action.RESET:
      return { ...initialState };
    default:
      throw new Error(`Invalid action type '${type}'`);
  }
}

export default function useLatestReducer(): [State, Dispatch<DispatchArgs>] {
  const [state, dispatch] = useReducer<Reducer<State, DispatchArgs>>(reducer, {
    ...initialState,
  });
  return [state, dispatch];
}
