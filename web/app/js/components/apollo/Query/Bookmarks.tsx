import { Query } from "react-apollo";
import query from "../graphql/query/bookmarks/search.graphql";
import { OffsetPagination } from "../../../types";
import { SearchParams } from "../../../types/bookmark";
import { SearchResults } from "../../../types/bookmark";

export interface Data {
  bookmarks: {
    search: SearchResults;
  };
}

export interface Variables {
  pagination: OffsetPagination;
  search?: SearchParams;
}

const variables = {
  pagination: {
    limit: 20
  }
};

export { query, variables };

class BookmarksQuery extends Query<Data, Variables> {
  static defaultProps = {
    query,
    variables
  };
}

export default BookmarksQuery;
