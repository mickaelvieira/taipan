import { Mutation } from "react-apollo";
import { UserBookmark } from "../../../types/bookmark";

interface Data {
  UpdateBookmark: UserBookmark;
}

interface Variables {
  url: string;
}

class UpdateBookmarkMutation extends Mutation<Data, Variables> {}

export default UpdateBookmarkMutation;
