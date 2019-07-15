import { Image } from "./image";
import { Event } from "./subscription";

export interface User {
  id: string;
  firstname: string;
  lastname: string;
  username: string;
  image: Image | null;
}

export interface UserEvent extends Event {
  item: User;
}
