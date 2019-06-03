import { Mutation } from "react-apollo";
import { Feed } from "../../../types/feed";
import mutation from "../../../services/apollo/mutation/feed.graphql";

interface Data {
  Feed: Feed;
}

interface Variables {
  url: string;
}

class FeedMutation extends Mutation<Data, Variables> {
  static defaultProps = {
    mutation
  };
}

export { mutation };

export default FeedMutation;
