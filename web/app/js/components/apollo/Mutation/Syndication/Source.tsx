import { Mutation } from "react-apollo";
import { Source } from "../../../../types/syndication";
import mutation from "../../graphql/mutation/syndication/source.graphql";

interface Data {
  syndication: {
    source: Source;
  };
}

interface Variables {
  url: string;
}

class SourceMutation extends Mutation<Data, Variables> {
  static defaultProps = {
    mutation
  };
}

export { mutation };

export default SourceMutation;
