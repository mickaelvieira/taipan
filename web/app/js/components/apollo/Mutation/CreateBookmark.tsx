import { Mutation } from "react-apollo";
import { Bookmark } from "../../../types/bookmark";
import mutation from "../../../services/apollo/mutation/create-bookmark.graphql";

interface Data {
  CreateBookmark: Bookmark;
}

interface Variables {
  url: string;
}

class CreateBookmarkMutation extends Mutation<Data, Variables> {
  static defaultProps = {
    mutation
  };
}

export { mutation };

export default CreateBookmarkMutation;
