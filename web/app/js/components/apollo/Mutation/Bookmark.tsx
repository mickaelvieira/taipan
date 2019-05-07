import { Mutation } from "react-apollo";
import { UserBookmark } from "../../../types/bookmark";
import mutation from "../../../services/apollo/mutation/bookmark.graphql";

interface Data {
  Bookmark: UserBookmark;
}

interface Variables {
  url: string;
}

class BookmarkMutation extends Mutation<Data, Variables> { }

export { mutation };

export default BookmarkMutation;
