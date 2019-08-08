import { Bookmark } from "../../../../types/bookmark";
import mutation from "../../graphql/mutation/bookmarks/favorite.graphql";

export interface Data {
  bookmarks: {
    favorite: Bookmark;
  };
}

export interface Variables {
  url: string;
}

export { mutation };
