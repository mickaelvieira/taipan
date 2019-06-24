import { Query } from "react-apollo";
import query from "../../../services/apollo/query/user.graphql";
import { User } from "../../../types/users";

export interface Data {
  user: User;
}

const variables = {};

export { query, variables };

class UserQuery extends Query<Data, {}> {
  static defaultProps = {
    query
  };
}

export default UserQuery;
