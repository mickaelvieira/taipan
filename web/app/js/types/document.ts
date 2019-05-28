import { Image } from "./image";

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
}
