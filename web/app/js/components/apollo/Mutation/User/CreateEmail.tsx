import { Mutation } from "react-apollo";
import { User } from "../../../../types/users";
import mutation from "../../graphql/mutation/user/create-email.graphql";

interface Data {
  users: {
    createEmail: User;
  };
}

interface Variables {
  email: string;
}

class CreateUserEmailMutation extends Mutation<Data, Variables> {
  static defaultProps = {
    mutation
  };
}

export { mutation };

export default CreateUserEmailMutation;
