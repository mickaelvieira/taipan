import { User } from "../../../../types/users";
import mutation from "../../graphql/mutation/user/profile.graphql";

export interface Data {
  users: {
    update: User;
  };
}

export interface Variables {
  user: {
    firstname: string;
    lastname: string;
    image: string;
  };
}

export { mutation };
