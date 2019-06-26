import { Mutation } from "react-apollo";
import { Source } from "../../../../types/syndication";
import enableSourceMutation from "../../../../services/apollo/mutation/syndication/enable-source.graphql";
import disableSourceMutation from "../../../../services/apollo/mutation/syndication/disable-source.graphql";

interface Data {
  syndication: {
    enable?: Source;
    disable?: Source;
  };
}

interface Variables {
  url: string;
}

class ChangeStatusMutation extends Mutation<Data, Variables> {
  static defaultProps = {};
}

export { enableSourceMutation, disableSourceMutation };

export default ChangeStatusMutation;
