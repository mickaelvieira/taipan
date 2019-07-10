import { Mutation } from "react-apollo";
import { User } from "../../../../types/users";
import mutation from "../../graphql/mutation/user/profile.graphql";

interface Data {
  users: {
    update: User;
  };
}

interface Variables {
  id: string;
  user: {
    firstname: string;
    lastname: string;
    image: string;
  };
}

class UserProfileMutation extends Mutation<Data, Variables> {
  static defaultProps = {
    mutation
  };
}

export { mutation };

export default UserProfileMutation;
