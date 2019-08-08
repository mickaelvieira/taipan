import { User } from "../../../../types/users";
import mutation from "../../graphql/mutation/user/primary-email.graphql";

export interface Data {
  users: {
    primaryEmail: User;
  };
}

export interface Variables {
  email: string;
}

export { mutation };
