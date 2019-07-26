import { Mutation } from "react-apollo";
import { Source } from "../../../../types/syndication";
import mutation from "../../graphql/mutation/syndication/title.graphql";

interface Data {
  syndication: {
    updateTitle: Source;
  };
}

interface Variables {
  url: string;
  title: string;
}

class UpdateSourceTitleMutation extends Mutation<Data, Variables> {
  static defaultProps = {
    mutation
  };
}

export { mutation };

export default UpdateSourceTitleMutation;
