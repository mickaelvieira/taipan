import { Mutation } from "react-apollo";
import { User } from "../../../../types/users";
import mutation from "../../graphql/mutation/user/primary-email.graphql";

interface Data {
  users: {
    primaryEmail: User;
  };
}

interface Variables {
  email: string;
}

class PrimaryUserEmailMutation extends Mutation<Data, Variables> {
  static defaultProps = {
    mutation
  };
}

export { mutation };

export default PrimaryUserEmailMutation;
