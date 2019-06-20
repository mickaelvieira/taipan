import dbFeed from "./feed";
import dbUser from "./user";
import dbBookmarks from "./bookmarks";
import { Bookmark } from "../../types/bookmark";

/* eslint @typescript-eslint/no-explicit-any: off */
export async function getDBState(): Promise<any> {
  const bkmks = await dbBookmarks.all();
  const feed = await dbFeed.all();
  const user = await dbUser.get();

  const bookmarks = new Map<string, Bookmark>();
  bkmks.forEach(item => {
    bookmarks.set(item.id, item);
  });

  return {
    bookmarks,
    feed,
    user
  };
}
