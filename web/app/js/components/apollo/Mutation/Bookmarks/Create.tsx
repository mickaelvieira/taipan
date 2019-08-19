import { Bookmark } from "../../../../types/bookmark";
import mutation from "../../graphql/mutation/bookmarks/create.graphql";

export interface Data {
  bookmarks: {
    create: Bookmark;
  };
}

export interface Variables {
  url: URL;
  isFavorite: boolean;
}

const variables = {
  isFavorite: false
};

export { mutation, variables };
