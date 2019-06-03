import { Mutation } from "react-apollo";
import { Bookmark } from "../../../types/bookmark";
import mutation from "../../../services/apollo/mutation/bookmark.graphql";

interface Data {
  Bookmark: Bookmark;
}

interface Variables {
  url: string;
}

class BookmarkMutation extends Mutation<Data, Variables> {
  static defaultProps = {
    mutation
  };
}

export { mutation };

export default BookmarkMutation;
