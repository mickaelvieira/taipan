import { User } from "../../../../types/users";
import mutation from "../../graphql/mutation/user/create-email.graphql";

export interface Data {
  users: {
    createEmail: User;
  };
}

export interface Variables {
  email: string;
}

export { mutation };
