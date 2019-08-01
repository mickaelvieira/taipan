import { Bookmark } from "./bookmark";
import { Document } from "./document";

export type SearchType = "bookmark" | "document";

export type SearchItem = Bookmark | Document;

export interface SearchResults {
  total: number;
  offset: number;
  limit: number;
  results: SearchItem[];
}

export interface SearchQueryData {
  [key: string]: {
    search: SearchResults;
  };
}
