import { Query } from "react-apollo";
import { UserBookmark } from "../../../types/bookmark";

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

class LatestBookmarksQuery extends Query<Data, Variables> {}

export default LatestBookmarksQuery;
