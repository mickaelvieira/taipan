import query from "../graphql/query/init.graphql";
import { AppInfo } from "../../../types/app";
import { User } from "../../../types/users";

export interface AppQueryData {
  app: {
    info: AppInfo | null;
  };
}

export interface UserQueryData {
  users: {
    loggedIn: User | null;
  };
}

export type InitQueryData = AppQueryData & UserQueryData;

const variables = {};

export { query, variables };
