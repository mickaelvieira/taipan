import { Query } from "react-apollo";
import query from "../../../services/apollo/query/users/logged-in.graphql";
import { AppInfo } from "../../../types/app";
import { User } from "../../../types/users";

export interface Data {
  app: {
    info: AppInfo;
  };
  users: {
    loggedIn: User;
  };
}

const variables = {};

export { query, variables };

class InitQuery extends Query<Data, {}> {
  static defaultProps = {
    query
  };
}

export default InitQuery;
