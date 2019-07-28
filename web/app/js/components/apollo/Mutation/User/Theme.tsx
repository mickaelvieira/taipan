import { Mutation } from "react-apollo";
import { User } from "../../../../types/users";
import mutation from "../../graphql/mutation/user/theme.graphql";

interface Data {
  users: {
    theme: User;
  };
}

interface Variables {
  id: string;
  theme: string;
}

class UserThemeMutation extends Mutation<Data, Variables> {
  static defaultProps = {
    mutation
  };
}

export { mutation };

export default UserThemeMutation;
