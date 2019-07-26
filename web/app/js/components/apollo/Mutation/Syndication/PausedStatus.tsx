import { Mutation } from "react-apollo";
import { Source } from "../../../../types/syndication";
import resumeMutation from "../../graphql/mutation/syndication/resume.graphql";
import pauseMutation from "../../graphql/mutation/syndication/pause.graphql";

interface Data {
  syndication: {
    pause?: Source;
    resume?: Source;
  };
}

interface Variables {
  url: string;
}

class SourcePausedStatusMutation extends Mutation<Data, Variables> {
  static defaultProps = {};
}

export { resumeMutation, pauseMutation };

export default SourcePausedStatusMutation;
