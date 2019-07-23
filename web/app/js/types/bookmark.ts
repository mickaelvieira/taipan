import { Image } from "./image";
import { Datetime } from "./scalars";

export interface Bookmark {
  id: string;
  url: string;
  lang?: string;
  charset?: string;
  title: string;
  description: string;
  image: Image | null;
  addedAt: Datetime;
  favoritedAt: Datetime;
  updatedAt: Datetime;
  isFavorite: boolean;
}
