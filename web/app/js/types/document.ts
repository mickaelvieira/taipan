import { Image } from "./image";
import { Source } from "./syndication";

export interface Document {
  id: string;
  url: string;
  lang?: string;
  charset?: string;
  title: string;
  description: string;
  image: Image | null;
  createdAt: string;
  updatedAt: string;
  syndication?: Source[];
}
