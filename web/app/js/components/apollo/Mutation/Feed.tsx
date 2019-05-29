import { Mutation } from "react-apollo";
import { Feed } from "../../../types/feed";
import mutation from "../../../services/apollo/mutation/feed.graphql";

interface Data {
  Feed: Feed;
}

interface Variables {
  url: string;
}

class FeedMutation extends Mutation<Data, Variables> {}

export { mutation };

export default FeedMutation;
