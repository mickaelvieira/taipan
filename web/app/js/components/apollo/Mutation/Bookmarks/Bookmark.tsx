import { Bookmark } from "../../../../types/bookmark";
import mutation from "../../graphql/mutation/bookmarks/bookmark.graphql";

export interface Data {
  bookmarks: {
    add: Bookmark;
  };
}

export interface Variables {
  url: URL;
  isFavorite: boolean;
  subscriptions?: string[];
}

export { mutation };
