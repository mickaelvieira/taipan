import { Tag } from "../../../../../types/syndication";
import mutation from "../../../graphql/mutation/syndication/tags/delete.graphql";

export interface Data {
  syndication: {
    deleteTag: Tag;
  };
}

export interface Variables {
  id: string;
}

export { mutation };
