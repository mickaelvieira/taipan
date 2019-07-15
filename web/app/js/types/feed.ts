import { Bookmark } from "./bookmark";
import { Document } from "./document";
import { Event, BusAction } from "./subscription";
import { CursorPagination } from ".";

export type FeedItem = Bookmark | Document;

export interface FeedResults {
  total: number;
  first: string;
  last: string;
  limit: number;
  results: FeedItem[];
}

export interface FeedEvent extends Event {
  item: FeedItem;
  action: FeedAction;
}

export interface FeedVariables {
  pagination: CursorPagination;
}

export interface FeedQueryData {
  feeds: {
    [key: string]: FeedResults;
  };
}

export interface FeedEventData {
  [key: string]: FeedEvent;
}

export type FeedAction = Extract<BusAction, "Add" | "Remove">;

export type FeedActions<T extends string> = {
  [index in T]: (result: FeedResults, item: FeedItem) => FeedResults;
};
