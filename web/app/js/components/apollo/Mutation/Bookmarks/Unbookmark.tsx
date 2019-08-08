import { Document } from "../../../../types/document";
import mutation from "../../graphql/mutation/bookmarks/unbookmark.graphql";

export interface Data {
  bookmarks: {
    remove: Document;
  };
}

export interface Variables {
  url: string;
}

export { mutation };
