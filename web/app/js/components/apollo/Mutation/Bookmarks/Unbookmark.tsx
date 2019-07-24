import { Mutation } from "react-apollo";
import { Document } from "../../../../types/document";
import mutation from "../../graphql/mutation/bookmarks/unbookmark.graphql";

interface Data {
  bookmarks: {
    remove: Document;
  };
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
