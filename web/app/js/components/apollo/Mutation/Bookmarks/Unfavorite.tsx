import { Mutation } from "react-apollo";
import { Bookmark } from "../../../../types/bookmark";
import mutation from "../../../../services/apollo/mutation/bookmarks/unfavorite.graphql";

interface Data {
  bookmarks: {
    unfavorite: Bookmark;
  };
}

interface Variables {
  url: string;
}

class UnfavoriteMutation extends Mutation<Data, Variables> {
  static defaultProps = {
    mutation
  };
}

export { mutation };

export default UnfavoriteMutation;
