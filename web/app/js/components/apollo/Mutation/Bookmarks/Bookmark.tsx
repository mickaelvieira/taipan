import { Mutation } from "react-apollo";
import { Bookmark } from "../../../../types/bookmark";
import mutation from "../../../../services/apollo/mutation/bookmarks/bookmark.graphql";

interface Data {
  bookmarks: {
    bookmark: Bookmark;
  };
}

interface Variables {
  url: string;
  isRead: boolean;
}

class BookmarkMutation extends Mutation<Data, Variables> {
  static defaultProps = {
    mutation
  };
}

export { mutation };

export default BookmarkMutation;
