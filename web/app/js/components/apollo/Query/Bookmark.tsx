import query from "../graphql/query/bookmarks/bookmark.graphql";
import { Bookmark } from "../../../types/bookmark";

export interface Data {
  bookmarks: {
    bookmark: Bookmark | null;
  };
}

export interface Variables {
  url: URL;
}

export { query };
