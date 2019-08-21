import query from "../graphql/query/users/logged-in.graphql";
import { User } from "../../../types/users";

export interface Data {
  users: {
    loggedIn: User | null;
  };
}

const variables = {};

export { query, variables };
