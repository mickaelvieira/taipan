import { Source } from "../../../../types/syndication";
import enableMutation from "../../graphql/mutation/syndication/enable.graphql";
import disableMutation from "../../graphql/mutation/syndication/disable.graphql";

export interface Data {
  syndication: {
    enable?: Source;
    disable?: Source;
  };
}

export interface Variables {
  url: string;
}

export { enableMutation, disableMutation };
