import { Query } from "react-apollo";
import { Bookmark } from "../../../types/bookmark";
import { Document } from "../../../types/document";
import query from "../../../services/apollo/query/feeds.graphql";

export interface Data {
  GetLatestBookmarks: {
    total: number;
    offset: number;
    limit: number;
    results: Bookmark[];
  };
  GetLatestDocuments: {
    total: number;
    offset: number;
    limit: number;
    results: Document[];
  };
}

export interface Variables {
  offset?: number;
  limit?: number;
}

const variables = {
  limit: 30
};

export { query, variables };

class LatestBookmarksQuery extends Query<Data, Variables> {}

export default LatestBookmarksQuery;
