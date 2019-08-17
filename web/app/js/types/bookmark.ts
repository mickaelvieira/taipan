import { Image } from "./image";

export interface Bookmark {
  id: string;
  url: URL;
  lang?: string;
  charset?: string;
  title: string;
  description: string;
  image: Image | null;
  addedAt: Date;
  favoritedAt: Date;
  updatedAt: Date;
  isFavorite: boolean;
}

export interface SearchResults {
  limit: number;
  total: number;
  offset: number;
  results: Bookmark[];
}

export interface SearchParams {
  terms: string[];
}
