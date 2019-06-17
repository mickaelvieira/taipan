import { Query } from "react-apollo";
import { Bookmark } from "../../../types/bookmark";
import { Document } from "../../../types/document";
import query from "../../../services/apollo/query/latest-news.graphql";
import { FeedQueryData, FeedVariables } from "../../../types/feed";

export type FeedItem = Bookmark | Document;

export interface CursorPagination {
  from?: string;
  to?: string;
  limit?: number;
}

const variables = {
  pagination: {
    limit: 10
  }
};
export { query, variables };

class LatestNewsQuery extends Query<FeedQueryData, FeedVariables> {
  static defaultProps = {
    query,
    variables,
    pollInterval: 0
  };
}

export default LatestNewsQuery;
