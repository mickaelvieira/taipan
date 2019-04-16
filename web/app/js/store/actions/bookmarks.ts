import API from "lib/api";
import { Dispatch } from "redux";
import {
  BookmarkActions,
  BookmarkResponseAction,
  BookmarkHistoryResponseAction
} from "store/actions/types";
import { CollectionResponse } from "collection/types";
import { Bookmark } from "types/bookmark";

export const updateBookmark = (
  info: CollectionResponse,
  history: CollectionResponse
): BookmarkResponseAction => ({
  type: BookmarkActions.UPSERT,
  payload: {
    info,
    history
  }
});

export const updateBookmarkHistory = (
  bookmark: Bookmark,
  history: CollectionResponse
): BookmarkHistoryResponseAction => ({
  type: BookmarkActions.HISTORY,
  payload: {
    bookmark,
    history
  }
});

export const fetchBookmark = (id: string) => async (dispatch: Dispatch) => {
  const results = await Promise.all([
    API.get(`/bookmark/${id}`),
    API.get(`/bookmark/${id}/history`)
  ]);

  dispatch(updateBookmark(...results));
  return results;
};

export const fetchBookmarkHistory = (bookmark: Bookmark) => async (
  dispatch: Dispatch
) => {
  const history = await API.get(`/bookmark/${bookmark.id}/history`);
  dispatch(updateBookmarkHistory(bookmark, history));
};
