import { Image } from "./image";
import { Event } from "./events";
import { ThemeName } from "../components/ui/themes";
import { Datetime } from "./scalars";

export interface User {
  id: string;
  firstname: string;
  lastname: string;
  username: string;
  emails: Email[];
  image: Image | null;
  theme?: ThemeName | null;
  createdAt: Datetime;
  updatedAt: Datetime;
}

export interface Email {
  id: string;
  value: string;
  isPrimary: boolean;
  isConfirmed: boolean;
  createdAt: Datetime;
  updatedAt: Datetime;
}

export interface UserEvent extends Event {
  item: User;
}
