import { Mutation } from "react-apollo";
import { Source } from "../../../../types/syndication";
import enableSourceMutation from "../../graphql/mutation/syndication/enable-source.graphql";
import disableSourceMutation from "../../graphql/mutation/syndication/disable-source.graphql";

interface Data {
  syndication: {
    enable?: Source;
    disable?: Source;
  };
}

interface Variables {
  url: string;
}

class ChangeSourceStatusMutation extends Mutation<Data, Variables> {
  static defaultProps = {};
}

export { enableSourceMutation, disableSourceMutation };

export default ChangeSourceStatusMutation;
