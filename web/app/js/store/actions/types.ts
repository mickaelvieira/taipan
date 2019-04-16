import { Action } from "redux";
import { CollectionResponse } from "collection/types";
import { Bookmark } from "types/bookmark";
import { RootState } from "store/reducer/default";

export enum IndexActions {
  FETCH = "@@index/FETCH",
  UPDATE = "@@index/UPDATE"
}

export enum FeedActions {
  APPEND = "@@feed/APPEND",
  PREPEND = "@@feed/PREPEND",
  UPSERT_ALL = "@@feed/UPSERT_ALL",
  DELETE = "@@feed/DELETE",
  REPLACE_ALL = "@@feed/REPLACE_ALL"
}

export enum BookmarkActions {
  UPSERT = "@@bookmark/UPSERT",
  HISTORY = "@@bookmark/HISTORY"
}

export enum UserActions {
  UPDATE = "@@user/UPDATE"
}

export interface ParamId {
  id: string;
}

export interface BookmarkHistoryPayload {
  history: CollectionResponse;
  bookmark: Bookmark;
}

export interface BookmarkDetails {
  history: CollectionResponse;
  info: CollectionResponse;
}

export interface BookmarkHistoryResponseAction {
  type: ReduxActionName;
  payload: BookmarkHistoryPayload;
}

export interface CollectionResponseAction extends Action {
  type: ReduxActionName;
  payload: CollectionResponse;
}

export interface DBStateResponseAction extends Action {
  type: ReduxActionName;
  payload: RootState;
}

export interface ParamIdAction extends Action {
  type: ReduxActionName;
  payload: ParamId;
}

export interface BookmarkResponseAction extends Action {
  type: ReduxActionName;
  payload: BookmarkDetails;
}

export type ReduxActionName =
  | IndexActions
  | FeedActions
  | UserActions
  | BookmarkActions;

export type ReduxActionWithPayload =
  | CollectionResponseAction
  | BookmarkResponseAction
  | BookmarkHistoryResponseAction
  | ParamIdAction
  | DBStateResponseAction;

export type ReduxAction = Action | ParamIdAction | CollectionResponseAction;

export type ReduxPayload =
  | ParamId
  | BookmarkHistoryPayload
  | CollectionResponse;
