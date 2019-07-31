import { Image } from "./image";
import { Event } from "./events";
import { ThemeName } from "../components/ui/themes";

export interface User {
  id: string;
  firstname: string;
  lastname: string;
  username: string;
  image: Image | null;
  theme?: ThemeName | null;
}

export interface UserEvent extends Event {
  item: User;
}
