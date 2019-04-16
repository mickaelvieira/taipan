import API from "lib/api";
import { ContentTypes } from "lib/api";
import { Action, Store, Dispatch } from "redux";
import { RootState } from "../reducer/default";
import { FeedActions, CollectionResponseAction } from "./types";
import { CollectionResponse } from "lib/collection/types";

export const CONNECTION_STATUS_CHANGE = "CONNECTION_STATUS_CHANGE";
export const SEARCH_RESULTS_UPDATE = "SEARCH_RESULTS_UPDATE";
export const SEARCH_TERMS_UPDATE = "SEARCH_TERMS_UPDATE";
export const HIDE_SEARCH_RESULTS = "HIDE_SEARCH_RESULTS";
export const SHOW_SEARCH_RESULTS = "SHOW_SEARCH_RESULTS";
export const SELECT_SEARCH_RESULT = "SELECT_SEARCH_RESULT";
export const PREPARE_SEARCH = "PREPARE_SEARCH";
export const LAUNCH_SEARCH = "LAUNCH_SEARCH";
export const WALL_ITEMS_FETCHING = "WALL_ITEMS_FETCHING";
export const WALL_ITEMS_UPDATE = "WALL_ITEMS_UPDATE";
export const WALL_ITEMS_ADD = "WALL_ITEMS_ADD";
export const WALL_ITEMS_DELETE = "WALL_ITEMS_DELETE";
export const WALL_ITEMS_REPLACE = "WALL_ITEMS_REPLACE";
export const WALL_ITEMS_PUSH = "WALL_ITEMS_PUSH";
export const WALL_ITEMS_PAGE = "WALL_ITEMS_PAGE";
export const SHOW_BOOKMARK = "SHOW_BOOKMARK";
export const UPDATE_BOOKMARK = "UPDATE_BOOKMARK";
export const HIDE_BOOKMARK = "HIDE_BOOKMARK";
export const UPDATE_FILTERS = "UPDATE_FILTERS";
export const UPDATE_USER = "UPDATE_USER";
export const HISTORY_PUSH_STATE = "HISTORY_PUSH_STATE";
export const HISTORY_REPLACE_STATE = "HISTORY_REPLACE_STATE";

export const historyPushState = (data, title: string, url: string) => ({
  type: HISTORY_PUSH_STATE,
  payload: {
    data,
    title,
    url
  }
});

export const historyReplaceState = (data, title: string, url: string) => ({
  type: HISTORY_REPLACE_STATE,
  payload: {
    data,
    title,
    url
  }
});

export const updateUser = (user: CollectionResponse) => ({
  type: UPDATE_USER,
  payload: {
    user
  }
});

export const updateConnectionStatus = (online: boolean) => ({
  type: CONNECTION_STATUS_CHANGE,
  payload: {
    online
  }
});

export const selectSearchResult = (selected: boolean) => ({
  type: SELECT_SEARCH_RESULT,
  payload: {
    selected
  }
});

export const hideSearchResults = () => ({
  type: HIDE_SEARCH_RESULTS,
  payload: {
    visible: false
  }
});

export const showSearchResults = () => ({
  type: SHOW_SEARCH_RESULTS,
  payload: {
    visible: true
  }
});

export const updateFilters = filter => ({
  type: UPDATE_FILTERS,
  payload: {
    filter
  }
});

export const updateSearchTerms = (terms: Array<string>) => ({
  type: SEARCH_TERMS_UPDATE,
  payload: {
    terms
  }
});

export const updateSearchResults = results => ({
  type: SEARCH_RESULTS_UPDATE,
  payload: {
    visible: true,
    fetching: false,
    results
  }
});

export const prepareSearch = () => ({
  type: PREPARE_SEARCH,
  payload: {
    selected: null,
    fetching: true,
    visible: true,
    results: []
  }
});

export const launchSearch = () => (dispatch, getState) => {
  const {
    search: { terms },
    feed: { links, filters }
  }: RootState = getState();
  const qTerms = terms
    .map(item => `terms[]=${encodeURIComponent(item)}`)
    .join("&");

  const qFilters = Object.keys(filters)
    .map(key => `filters[${key}]=${encodeURIComponent(filters[key])}`)
    .join("&");

  const query = [qTerms, qFilters].join("&");

  return API.get(links.search + "?" + query)
    .then(({ collection: { items = [] } }) => {
      dispatch(updateSearchResults(items));
    })
    .catch(error => {
      dispatch(updateSearchResults([]));
      throw error;
    });
};

export const markWallAsFetching = () => ({
  type: WALL_ITEMS_FETCHING,
  payload: {
    fetching: true
  }
});

export const addWallItems = (
  data: CollectionResponse
): CollectionResponseAction => ({
  type: FeedActions.PREPEND,
  payload: data
});

export const replaceWallItems = (
  data: CollectionResponse
): CollectionResponseAction => ({
  type: FeedActions.REPLACE_ALL,
  payload: data
});

export const pushWallItems = (
  data: CollectionResponse
): CollectionResponseAction => ({
  type: FeedActions.APPEND,
  payload: data
});

export const removeWallItems = id => ({
  type: WALL_ITEMS_DELETE,
  payload: {
    id
  }
});

export const updateWallItems = (
  data: CollectionResponse
): CollectionResponseAction => ({
  type: FeedActions.UPSERT_ALL,
  payload: data
});

/**
 * @returns {Function}
 */
export const refreshBookmark = ({ links }) => (dispatch, getState) => {
  return API.post(links.about).then(json => {
    dispatch(updateWallItems(json));
  });
};

/**
 * @returns {Function}
 */
export const removeBookmark = ({ id, links }) => (dispatch, getState) => {
  return API.delete(links.self, ContentTypes.HTML).then(() =>
    dispatch(removeWallItems(id))
  );
};

export const markAsRead = ({ links }) => (dispatch, getState) => {
  return API.put(links.read).then(json => {
    console.log(json);
    dispatch(updateWallItems(json));
  });
};

export const markAsUnread = ({ links }) => (dispatch, getState) => {
  return API.put(links.unread, null, ContentTypes.HTML).then(() =>
    dispatch(updateWallItems(json))
  );
};

export const fetchItems = () => (dispatch, getState) => {
  return API.get(getState().feed.links.first).then(json =>
    dispatch(replaceWallItems(json))
  );
};

export const fetchItemsWithFilters = () => (dispatch, getState) => {
  const { links, filters }: RootState = getState().feed;
  const query = Object.keys(filters)
    .map(key => `filters[${key}]=${encodeURIComponent(filters[key])}`)
    .join("&");

  const url = `${links.search}?${query}`;

  return API.get(url).then(json => dispatch(replaceWallItems(json)));
};

export const fetchIndex = () => (dispatch, getState) => {
  return API.get("/").then(json => dispatch(replaceWallItems(json)));
};

export const fetchUser = () => (dispatch, getState) => {
  return API.get(getState().feed.links.user).then(json =>
    dispatch(updateUser(json))
  );
};

export const fetchNextPageItems = () => (dispatch, getState) => {
  return API.get(getState().feed.links.next).then(json =>
    dispatch(pushWallItems(json))
  );
};

export const showBookmark = () => ({
  type: SHOW_BOOKMARK,
  payload: {
    visible: true
  }
});

export const updateBookmark = (html, info, history) => ({
  type: UPDATE_BOOKMARK,
  payload: {
    html,
    info,
    history,
    visible: true
  }
});

export const hideBookmark = () => ({
  type: HIDE_BOOKMARK,
  payload: {
    visible: false
  }
});

export const fetchBookmark = (id: string) => (dispatch, getState) => {
  return Promise.all([
    API.get(`/bookmark/${id}`, ContentTypes.HTML),
    API.get(`/bookmark/${id}`),
    API.get(`/bookmark/${id}/history`)
  ]).then(([response, info, history]) => {
    response.text().then((html: string) => {
      dispatch(updateBookmark(html, info, history));
    });
  });
};
