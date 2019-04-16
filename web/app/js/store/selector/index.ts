import { isEqual } from "lodash";
import { createSelectorCreator, defaultMemoize } from "reselect";
import { RootState } from "store/reducer/default";
import { Bookmark } from "types/bookmark";

const createDeepEqualSelector = createSelectorCreator(defaultMemoize, isEqual);

export const selectFeedBookmarks = createDeepEqualSelector(
  (state: RootState) => state.feed.items,
  (state: RootState) => state.index.bookmarks,
  (ids: string[], bookmarks: Map<string, Bookmark>) =>
    ids.map(id => bookmarks.get(id)).filter(bookmark => bookmark)
);

export const getCurrentIndex = (id?: string, items?: Bookmark[]) => {
  let current = 0;
  if (id && items && items.length > 0) {
    current = items.findIndex(item => item.id === id);
    if (current < 0) {
      current = 0;
    }
  }

  return current;
};
