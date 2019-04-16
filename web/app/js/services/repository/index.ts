import dbFeed from "./feed";
import dbUser from "./user";
import dbBookmarks from "./bookmarks";
import defaultState, { RootState } from "store/reducer/default";
import { Bookmark } from "types/bookmark";

export async function getDBState(): Promise<RootState> {
  const bkmks = await dbBookmarks.all();
  const feed = await dbFeed.all();
  const user = await dbUser.get();

  const bookmarks = new Map<string, Bookmark>();
  bkmks.forEach(item => {
    bookmarks.set(item.id, item);
  });

  return {
    ...defaultState,
    index: {
      ...defaultState.index,
      bookmarks
    },
    feed: {
      ...defaultState.feed,
      items: feed.results
    },
    user
  };
}
