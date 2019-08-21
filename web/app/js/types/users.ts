import { Image } from "./image";
import { Event } from "./events";
import { ThemeName } from "../components/ui/themes";

export interface User {
  id: string;
  firstname: string;
  lastname: string;
  emails: Email[];
  image: Image | null;
  stats?: Stats;
  theme?: ThemeName | null;
  createdAt: Date;
  updatedAt: Date;
}

export interface Stats {
  bookmarks: number;
  favorites: number;
  readingList: number;
  subscriptions: number;
}

export interface Email {
  id: string;
  value: string;
  isPrimary: boolean;
  isConfirmed: boolean;
  createdAt: Date;
  updatedAt: Date;
}

export interface UserEvent extends Event {
  item: User;
}
