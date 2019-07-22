import { Mutation } from "react-apollo";
import { Bookmark } from "../../../../types/bookmark";
import mutation from "../../graphql/mutation/bookmarks/bookmark.graphql";

interface Data {
  bookmarks: {
    add: Bookmark;
  };
}

interface Variables {
  url: string;
  isFavorite: boolean;
  subscriptions?: string[];
}

class BookmarkMutation extends Mutation<Data, Variables> {
  static defaultProps = {
    mutation
  };
}

export { mutation };

export default BookmarkMutation;
