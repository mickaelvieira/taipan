import { Mutation } from "react-apollo";
import { UserBookmark } from "../../../types/bookmark";
import mutation from "../../../services/apollo/mutation/update-bookmark.graphql";

interface Data {
  UpdateBookmark: UserBookmark;
}

interface Variables {
  url: string;
}

class UpdateBookmarkMutation extends Mutation<Data, Variables> {}

export { mutation };

export default UpdateBookmarkMutation;
