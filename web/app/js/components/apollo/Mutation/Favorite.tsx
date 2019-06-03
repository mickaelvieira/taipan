import { Mutation } from "react-apollo";
import { Bookmark } from "../../../types/bookmark";
import mutation from "../../../services/apollo/mutation/favorite.graphql";

interface Data {
  ChangeBookmarkReadStatus: Bookmark;
}

interface Variables {
  url: string;
  isFavorite: boolean;
}

class FavoriteMutation extends Mutation<Data, Variables> {
  static defaultProps = {
    mutation
  };
}

export { mutation };

export default FavoriteMutation;
