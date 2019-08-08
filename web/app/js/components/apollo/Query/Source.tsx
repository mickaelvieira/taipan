import query from "../graphql/query/syndication/source.graphql";
import { Source } from "../../../types/syndication";

export interface Data {
  syndication: {
    source: Source | null;
  };
}

export interface Variables {
  url: string;
}

export { query };
