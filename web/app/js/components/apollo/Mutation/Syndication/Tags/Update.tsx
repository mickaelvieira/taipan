import { Tag } from "../../../../../types/syndication";
import mutation from "../../../graphql/mutation/syndication/tags/update.graphql";

export interface Data {
  syndication: {
    updateTag: Tag;
  };
}

export interface Variables {
  id: string;
  label: string;
}

export { mutation };
