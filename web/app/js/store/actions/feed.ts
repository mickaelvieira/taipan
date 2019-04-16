import { Dispatch } from "redux";
import API, { ContentTypes } from "lib/api";
import { FeedActions, CollectionResponseAction, ParamIdAction } from "./types";
import { CollectionResponse } from "collection/types";
import { Bookmark } from "types/bookmark";
import { RootState } from "store/reducer/default";

interface GetState {
  (): RootState;
}

export const addFeedItems = (
  data: CollectionResponse
): CollectionResponseAction => ({
  type: FeedActions.PREPEND,
  payload: data
});

export const replaceFeedItems = (
  data: CollectionResponse
): CollectionResponseAction => ({
  type: FeedActions.REPLACE_ALL,
  payload: data
});

export const pushFeedItems = (
  data: CollectionResponse
): CollectionResponseAction => ({
  type: FeedActions.APPEND,
  payload: data
});

export const removeFeedItems = (id: string): ParamIdAction => ({
  type: FeedActions.DELETE,
  payload: {
    id
  }
});

export const updateFeedItems = (
  data: CollectionResponse
): CollectionResponseAction => ({
  type: FeedActions.UPSERT_ALL,
  payload: data
});

export const refreshBookmark = ({ links }: Bookmark) => async (
  dispatch: Dispatch
) => {
  const json = await API.post(links.about);
  return dispatch(updateFeedItems(json));
};

export const removeBookmark = ({ id, links }: Bookmark) => async (
  dispatch: Dispatch
) => {
  await API.delete(links.self, ContentTypes.HTML);
  return dispatch(removeFeedItems(id));
};

export const markAsRead = ({ links }: Bookmark) => async (
  dispatch: Dispatch
) => {
  const json = await API.put(links.read);
  return dispatch(updateFeedItems(json));
};

export const markAsUnread = ({ links }: Bookmark) => async (
  dispatch: Dispatch
) => {
  const json = await API.put(links.unread, null, ContentTypes.HTML);
  return dispatch(updateFeedItems(json));
};

export const fetchItems = () => async (
  dispatch: Dispatch,
  getState: GetState
) => {
  const state = getState();
  const {
    index: { links }
  } = state;

  if (typeof links.first === "undefined") {
    throw new Error("first link is not defined");
  }

  const json = await API.get(links.first);
  return dispatch(replaceFeedItems(json));
};

export const fetchNextPageItems = () => async (
  dispatch: Dispatch,
  getState: GetState
) => {
  const state = getState();
  const {
    feed: { links }
  } = state;

  if (typeof links.next === "undefined") {
    throw new Error("next link is not defined");
  }

  const json = await API.get(links.next);
  return dispatch(pushFeedItems(json));
};
