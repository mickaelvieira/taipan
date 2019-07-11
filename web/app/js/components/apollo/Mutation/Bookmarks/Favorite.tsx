import { Mutation } from "react-apollo";
import { Bookmark } from "../../../../types/bookmark";
import mutation from "../../graphql/mutation/bookmarks/favorite.graphql";

interface Data {
  bookmarks: {
    favorite: Bookmark;
  };
}

interface Variables {
  url: string;
}

class FavoriteMutation extends Mutation<Data, Variables> {
  static defaultProps = {
    mutation
  };
}

export { mutation };

export default FavoriteMutation;
