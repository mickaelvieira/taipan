import { Image } from "./image";
import { Source } from "./syndication";
import { Datetime } from "./scalars";

export interface Document {
  id: string;
  url: string;
  lang?: string;
  charset?: string;
  title: string;
  description: string;
  image: Image | null;
  createdAt: Datetime;
  updatedAt: Datetime;
  syndication?: Source[];
}

export interface SearchResults {
  limit: number;
  total: number;
  offset: number;
  results: Document[];
}

export interface SearchParams {
  terms: string[];
}
