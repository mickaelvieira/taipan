import { Mutation } from "react-apollo";
import { Bookmark } from "../../../../types/bookmark";
import mutation from "../../graphql/mutation/bookmarks/create.graphql";

interface Data {
  bookmarks: {
    create: Bookmark;
  };
}

interface Variables {
  url: string;
  withFeeds: boolean;
}

const variables = {
  withFeeds: false
};

class CreateBookmarkMutation extends Mutation<Data, Variables> {
  static defaultProps = {
    mutation,
    variables
  };
}

export { mutation, variables };

export default CreateBookmarkMutation;
