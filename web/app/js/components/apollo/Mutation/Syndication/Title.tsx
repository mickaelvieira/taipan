import { Source } from "../../../../types/syndication";
import mutation from "../../graphql/mutation/syndication/title.graphql";

export interface Data {
  syndication: {
    updateTitle: Source;
  };
}

export interface Variables {
  url: string;
  title: string;
}

export { mutation };
