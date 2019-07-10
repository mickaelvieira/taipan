import { Image } from "./image";

export interface User {
  id: string;
  firstname: string;
  lastname: string;
  username: string;
  image: Image | null;
}
