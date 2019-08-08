import { User } from "../../../../types/users";
import mutation from "../../graphql/mutation/user/delete-email.graphql";

export interface Data {
  users: {
    deleteEmail: User;
  };
}

export interface Variables {
  email: string;
}

export { mutation };
