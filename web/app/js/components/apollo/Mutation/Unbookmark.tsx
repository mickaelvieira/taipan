import { Mutation } from "react-apollo";
import { Document } from "../../../types/document";
import mutation from "../../../services/apollo/mutation/unbookmark.graphql";

interface Data {
  UnbookmarkMutation: Document;
}

interface Variables {
  url: string;
}

class UnbookmarkMutation extends Mutation<Data, Variables> {}

export { mutation };

export default UnbookmarkMutation;
