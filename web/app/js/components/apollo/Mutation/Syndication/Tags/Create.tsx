import { Tag } from "../../../../../types/syndication";
import mutation from "../../../graphql/mutation/syndication/tags/create.graphql";

export interface Data {
  syndication: {
    createTag: Tag;
  };
}

export interface Variables {
  label: string;
}

export { mutation };
