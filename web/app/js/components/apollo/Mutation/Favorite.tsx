import { Mutation } from "react-apollo";
import { Bookmark } from "../../../types/bookmark";
import mutation from "../../../services/apollo/mutation/favorite.graphql";

interface Data {
  Bookmark: Bookmark;
}

interface Variables {
  url: string;
  isFavorite: boolean;
}

class FavoriteMutation extends Mutation<Data, Variables> {}

export { mutation };

export default FavoriteMutation;
