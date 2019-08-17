import { User } from "../../../../types/users";
import mutation from "../../graphql/mutation/user/theme.graphql";

export interface Data {
  users: {
    theme: User;
  };
}

export interface Variables {
  theme: string;
}

export { mutation };
