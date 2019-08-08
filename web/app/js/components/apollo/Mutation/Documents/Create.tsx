import { Document } from "../../../../types/document";
import mutation from "../../graphql/mutation/documents/create.graphql";

export interface Data {
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

export { mutation, variables };
