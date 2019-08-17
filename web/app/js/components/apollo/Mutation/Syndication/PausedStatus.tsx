import { Source } from "../../../../types/syndication";
import resumeMutation from "../../graphql/mutation/syndication/resume.graphql";
import pauseMutation from "../../graphql/mutation/syndication/pause.graphql";

export interface Data {
  syndication: {
    pause?: Source;
    resume?: Source;
  };
}

export interface Variables {
  url: URL;
}

export { resumeMutation, pauseMutation };
