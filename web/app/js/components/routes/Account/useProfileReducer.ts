import { useReducer, Reducer, Dispatch } from "react";
import { User } from "../../../types/users";

interface State {
  firstname: string;
  lastname: string;
  file: File | null;
  scale: number;
}

type Payload = string | number | File | null;

export enum Action {
  FIRSTNAME = "firstname",
  LASTNAME = "lastname",
  AVATAR = "avatar",
  SCALE = "scale"
}

function reducer(state: State, [type, payload]: [Action, Payload]): State {
  switch (type) {
    case Action.FIRSTNAME:
      return {
        ...state,
        firstname: payload as string
      };
    case Action.LASTNAME:
      return {
        ...state,
        lastname: payload as string
      };
    case Action.AVATAR:
      return {
        ...state,
        file: payload as File
      };
    case Action.SCALE:
      return {
        ...state,
        scale: payload as number
      };
    default:
      throw new Error(`Invalid action type '${type}'`);
  }
}

function getInitialState(user: User | null): State {
  const scale = 1;
  const file = null;
  let firstname = "";
  let lastname = "";

  if (user) {
    firstname = user.firstname;
    lastname = user.lastname;
  }

  return {
    firstname,
    lastname,
    scale,
    file
  };
}

type ProfileReducer = Reducer<State, [Action, Payload]>;

export default function useProfileReducer(
  user: User
): [State, Dispatch<[Action, Payload]>] {
  const [state, dispatch] = useReducer<ProfileReducer, User | null>(
    reducer,
    user,
    getInitialState
  );

  return [state, dispatch];
}
