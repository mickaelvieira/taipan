import { IndexState, defaultState as index } from "store/reducer/index";
import { FeedState, defaultState as feed } from "store/reducer/feed";
import { UserState, defaultState as user } from "store/reducer/user";

export interface RootState {
  index: IndexState;
  feed: FeedState;
  user: UserState;
}

export default {
  index,
  feed,
  user
};
