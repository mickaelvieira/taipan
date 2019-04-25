import { Mutation } from "react-apollo";
import { UserBookmark } from "../../../types/bookmark";

interface Data {
  CreateBookmark: UserBookmark;
}

interface Variables {
  url: string;
}

class CreateBookmarkMutation extends Mutation<Data, Variables> {}

export default CreateBookmarkMutation;
