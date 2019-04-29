import { Mutation } from "react-apollo";
import { UserBookmark } from "../../../types/bookmark";
import mutation from "../../../services/apollo/mutation/create-bookmark.graphql";

interface Data {
  CreateBookmark: UserBookmark;
}

interface Variables {
  url: string;
}

class CreateBookmarkMutation extends Mutation<Data, Variables> {}

export { mutation };

export default CreateBookmarkMutation;
