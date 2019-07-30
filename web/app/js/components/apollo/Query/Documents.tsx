import { Query } from "react-apollo";
import query from "../graphql/query/documents/search.graphql";
import { OffsetPagination } from "../../../types";
import { SearchParams } from "../../../types/document";
import { SearchResults } from "../../../types/document";

export interface Data {
  documents: {
    search: SearchResults;
  };
}

export interface Variables {
  pagination: OffsetPagination;
  search?: SearchParams;
}

const variables = {
  pagination: {
    limit: 50
  }
};

export { query, variables };

class DocumentsQuery extends Query<Data, Variables> {
  static defaultProps = {
    query,
    variables
  };
}

export default DocumentsQuery;
