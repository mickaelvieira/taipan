import { User } from "../../../../types/users";
import mutation from "../../graphql/mutation/user/password.graphql";

export interface Data {
  users: {
    password: User;
  };
}

export interface Variables {
  old: string;
  new: string;
}

export { mutation };
