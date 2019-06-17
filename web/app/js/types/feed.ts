import { Bookmark } from "./bookmark";
import { Document } from "./document";

export interface Feed {
  id: string;
  url: string;
  title: string;
  type: string;
  status: string;
  createdAt: string;
  updatedAt: string;
  parsedAt: string;
}

export type FeedItem = Bookmark | Document;

export interface CursorPagination {
  from?: string;
  to?: string;
  limit?: number;
}

export interface FeedQueryResult {
  total: number;
  first: string;
  last: string;
  limit: number;
  results: FeedItem[];
}

export interface BookmarkEvent {
  id: string;
  bookmark: Bookmark;
  action: FeedAction;
  topic: FeedTopic;
}

export interface FeedVariables {
  pagination: CursorPagination;
}

export interface FeedData {
  [key: string]: FeedQueryResult;
}

export type FeedTopic = "News" | "Favorites" | "ReadingList";
export type FeedAction = "Add" | "Remove";

export type FeedActions<T extends string> = {
  [index in T]: (result: FeedQueryResult, item: FeedItem) => FeedQueryResult;
};
