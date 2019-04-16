import API from "lib/api";
import { Dispatch } from "redux";
import { UserActions } from "./types";
import { CollectionResponse } from "collection/types";
import { RootState } from "store/reducer/default";

interface GetState {
  (): RootState;
}

export const updateUser = (data: CollectionResponse) => ({
  type: UserActions.UPDATE,
  payload: data
});

export const fetchUser = () => async (
  dispatch: Dispatch,
  getState: GetState
) => {
  const state = getState();
  const {
    index: { links }
  } = state;

  if (typeof links.user === "undefined") {
    throw new Error("user link is not defined");
  }

  const json = await API.get(links.user);
  return dispatch(updateUser(json));
};
