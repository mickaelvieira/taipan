import { Image } from "./image";
import { Source } from "./syndication";

export interface Document {
  id: string;
  url: URL;
  lang?: string;
  charset?: string;
  title: string;
  description: string;
  image: Image | null;
  createdAt: Date;
  updatedAt: Date;
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
