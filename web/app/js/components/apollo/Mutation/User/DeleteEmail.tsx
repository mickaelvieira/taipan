import { Mutation } from "react-apollo";
import { User } from "../../../../types/users";
import mutation from "../../graphql/mutation/user/delete-email.graphql";

interface Data {
  users: {
    deleteEmail: User;
  };
}

interface Variables {
  email: string;
}

class DeleteUserEmailMutation extends Mutation<Data, Variables> {
  static defaultProps = {
    mutation
  };
}

export { mutation };

export default DeleteUserEmailMutation;
