import { Query } from "react-apollo";
import query from "../graphql/query/bookmarks/bookmark.graphql";
import { Bookmark } from "../../../types/bookmark";

export interface Data {
  bookmarks: {
    bookmark: Bookmark | null;
  };
}

export interface Variables {
  url: string;
}

export { query };

class BookmarkQuery extends Query<Data, Variables> {
  static defaultProps = {
    query
  };
}

export default BookmarkQuery;
