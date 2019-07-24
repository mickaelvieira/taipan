import { Bookmark } from "./bookmark";
import { Document } from "./document";
import { Event } from "./events";
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

export type FeedName = "news" | "readinglist" | "favorites";
