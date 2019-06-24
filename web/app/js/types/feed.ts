import { Bookmark } from "./bookmark";
import { Document } from "./document";

export type FeedItem = Bookmark | Document;

export interface CursorPagination {
  from?: string;
  to?: string;
  limit?: number;
}

export interface FeedResults {
  total: number;
  first: string;
  last: string;
  limit: number;
  results: FeedItem[];
}

export interface FeedEvent {
  id: string;
  item: FeedItem;
  action: FeedAction;
  topic: FeedTopic;
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

export type FeedTopic = "News" | "Favorites" | "ReadingList";
export type FeedAction = "Add" | "Remove";

export type FeedActions<T extends string> = {
  [index in T]: (result: FeedResults, item: FeedItem) => FeedResults;
};
