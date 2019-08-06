import { Query } from "react-apollo";
import query from "../graphql/query/init.graphql";
import { AppInfo } from "../../../types/app";
import { User } from "../../../types/users";

export interface AppQueryData {
  app: {
    info: AppInfo;
  };
}

export interface UserQueryData {
  users: {
    loggedIn: User | null;
  };
}

export type Data = AppQueryData & UserQueryData;

const variables = {};

export { query, variables };

class InitQuery extends Query<Data, {}> {
  static defaultProps = {
    query
  };
}

export default InitQuery;
