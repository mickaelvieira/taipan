import { Mutation } from "react-apollo";
import { Document } from "../../../types/document";
import mutation from "../../../services/apollo/mutation/unbookmark.graphql";

interface Data {
  Unbookmark: Document;
}

interface Variables {
  url: string;
}

class UnbookmarkMutation extends Mutation<Data, Variables> {
  static defaultProps = {
    mutation
  };
}

export { mutation };

export default UnbookmarkMutation;
