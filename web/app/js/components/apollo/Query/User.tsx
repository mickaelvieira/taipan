import { Query } from "react-apollo";
import query from "../../../services/apollo/query/users/logged-in.graphql";
import { User } from "../../../types/users";

export interface Data {
  users: {
    loggedIn: User;
  };
}

const variables = {};

export { query, variables };

class UserQuery extends Query<Data, {}> {
  static defaultProps = {
    query
  };
}

export default UserQuery;
