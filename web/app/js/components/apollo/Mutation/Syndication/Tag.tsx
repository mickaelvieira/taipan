import { Source } from "../../../../types/syndication";
import tagMutation from "../../graphql/mutation/syndication/tag.graphql";
import untagMutation from "../../graphql/mutation/syndication/untag.graphql";

export interface Data {
  syndication: {
    tag?: Source;
    untag?: Source;
  };
}

export interface Variables {
  sourceId: string;
  tagId: string;
}

export { tagMutation, untagMutation };
