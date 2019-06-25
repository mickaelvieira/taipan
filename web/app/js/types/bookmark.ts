import { Image } from "./image";

export interface Bookmark {
  id: string;
  url: string;
  lang?: string;
  charset?: string;
  title: string;
  description: string;
  image: Image | null;
  addedAt: string;
  updatedAt: string;
  isFavorite: boolean;
}
