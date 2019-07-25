import { Query } from "react-apollo";
import query from "../graphql/query/syndication/source.graphql";
import { Source } from "../../../types/syndication";

export interface Data {
  syndication: {
    source: Source;
  };
}

export interface Variables {
  url: string;
}

export { query };

class SourceQuery extends Query<Data, Variables> {
  static defaultProps = {
    query
  };
}

export default SourceQuery;
