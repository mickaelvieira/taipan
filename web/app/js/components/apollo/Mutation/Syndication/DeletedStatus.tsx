import { Mutation } from "react-apollo";
import { Source } from "../../../../types/syndication";
import enableMutation from "../../graphql/mutation/syndication/enable.graphql";
import disableMutation from "../../graphql/mutation/syndication/disable.graphql";

interface Data {
  syndication: {
    enable?: Source;
    disable?: Source;
  };
}

interface Variables {
  url: string;
}

class SourceDeletedStatusMutation extends Mutation<Data, Variables> {
  static defaultProps = {};
}

export { enableMutation, disableMutation };

export default SourceDeletedStatusMutation;
