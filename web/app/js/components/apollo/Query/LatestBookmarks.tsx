import { Query } from "react-apollo";
import { UserBookmark } from "../../../types/bookmark";
import query from "../../../services/apollo/query/latest-bookmarks.graphql";

interface Data {
  GetLatestBookmarks: {
    total: number;
    offset: number;
    limit: number;
    results: UserBookmark[];
  };
}

interface Variables {
  offset?: number;
  limit?: number;
}

const variables = {
  limit: 30
};

export { query, variables };

class LatestBookmarksQuery extends Query<Data, Variables> {}

export default LatestBookmarksQuery;
