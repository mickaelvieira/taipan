import { Query } from "react-apollo";
import { Bookmark } from "../../../types/bookmark";
import { Document } from "../../../types/document";
import query from "../../../services/apollo/query/latest-news.graphql";

export type FeedItem = Bookmark | Document;

export interface CursorPagination {
  from?: string;
  to?: string;
  limit?: number;
}

export interface NewsResult {
  results: Document[];
}

export interface Variables {
  pagination: CursorPagination;
}

export interface Data {
  [key: string]: NewsResult;
}

const variables = {
  pagination: {
    limit: 10
  }
};
export { query, variables };

class LatestNewsQuery extends Query<Data, Variables> {
  static defaultProps = {
    query,
    variables,
    pollInterval: 0
  };
}

export default LatestNewsQuery;
