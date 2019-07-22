import { Mutation } from "react-apollo";
import { Document } from "../../../../types/document";
import mutation from "../../graphql/mutation/documents/create.graphql";

interface Data {
  documents: {
    create: Document;
  };
}

export interface Variables {
  url: string;
}

const variables = {
  isFavorite: false,
  withFeeds: true
};

class CreateDocumentMutation extends Mutation<Data, Variables> {
  static defaultProps = {
    mutation,
    variables
  };
}

export { mutation, variables };

export default CreateDocumentMutation;
