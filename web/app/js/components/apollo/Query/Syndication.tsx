import { Query } from "react-apollo";
import query from "../../../services/apollo/query/syndication/sources.graphql";
import { SyndicationResults } from "../../../types/syndication";

export interface Data {
  syndication: {
    sources: SyndicationResults;
  };
}

const variables = {
  pagination: {
    limit: 50
  }
};

export { query, variables };

class SyndicationQuery extends Query<Data, {}> {
  static defaultProps = {
    query,
    variables
  };
}

export default SyndicationQuery;
