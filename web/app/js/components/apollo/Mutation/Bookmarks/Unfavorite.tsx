import { Bookmark } from "../../../../types/bookmark";
import mutation from "../../graphql/mutation/bookmarks/unfavorite.graphql";

export interface Data {
  bookmarks: {
    unfavorite: Bookmark;
  };
}

export interface Variables {
  url: URL;
}

export { mutation };
