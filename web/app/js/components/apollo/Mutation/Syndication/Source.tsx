import { Source } from "../../../../types/syndication";
import mutation from "../../graphql/mutation/syndication/source.graphql";

export interface Data {
  syndication: {
    source: Source;
  };
}

export interface Variables {
  url: string;
  tags: string[];
}

export { mutation };
