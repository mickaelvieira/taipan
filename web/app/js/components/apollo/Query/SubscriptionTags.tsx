import query from "../graphql/query/subscriptions/tags.graphql";
import { TagResults } from "../../../types/subscription";

export interface Data {
  subscriptions: {
    tags: TagResults;
  };
}

const variables = {};

export { query, variables };
