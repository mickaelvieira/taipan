import API from "lib/api";
import { IndexActions, CollectionResponseAction } from "./types";
import { CollectionResponse } from "collection/types";
import { Dispatch } from "redux";

export const fetchIndex = () => async (dispatch: Dispatch) => {
  const json = await API.get("/");
  return dispatch(updateIndex(json));
};

export const updateIndex = (
  data: CollectionResponse
): CollectionResponseAction => ({
  type: IndexActions.UPDATE,
  payload: data
});
