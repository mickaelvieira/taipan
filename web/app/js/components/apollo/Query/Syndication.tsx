import { Query } from "react-apollo";
import query from "../../../services/apollo/query/syndication/sources.graphql";
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

class SyndicationQuery extends Query<Data, Variables> {
  static defaultProps = {
    query,
    variables
  };
}

export default SyndicationQuery;
