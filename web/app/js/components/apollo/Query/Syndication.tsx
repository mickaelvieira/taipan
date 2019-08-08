import query from "../graphql/query/syndication/sources.graphql";
import { OffsetPagination } from "../../../types";
import { SyndicationResults, SearchParams } from "../../../types/syndication";

export interface Data {
  syndication: {
    sources: SyndicationResults;
  };
}

export interface Variables {
  pagination: OffsetPagination;
  search: SearchParams;
}

const variables = {
  pagination: {
    limit: 50
  },
  search: {
    isPaused: false
  }
};

export { query, variables };
