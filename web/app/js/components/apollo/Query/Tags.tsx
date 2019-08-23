import query from "../graphql/query/syndication/tags.graphql";
import { SyndicationTagResults } from "../../../types/syndication";

export interface Data {
  syndication: {
    tags: SyndicationTagResults;
  };
}

const variables = {};

export { query, variables };
