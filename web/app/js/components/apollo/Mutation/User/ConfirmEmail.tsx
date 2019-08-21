import { User } from "../../../../types/users";
import mutation from "../../graphql/mutation/user/send-confirm-email.graphql";

export interface Data {
  users: {
    sendConfirmationEmail: User;
  };
}

export interface Variables {
  email: string;
}

export { mutation };
