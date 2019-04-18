import {
  FeedActions,
  IndexActions,
  BookmarkActions,
  BookmarkDetails,
  ReduxActionWithPayload,
  BookmarkHistoryPayload
} from "store/actions/types";
import { Index, IndexLinks } from "types/index";
import { Bookmarks, Bookmark, BookmarkHistory } from "types/bookmark";
import { parseItem, parseCollection } from "collection";
import { CollectionResponse } from "collection/types";
import db from "services/repository/bookmarks";

export interface IndexState {
  bookmarks: Map<string, Bookmark>;
  links: IndexLinks;
}

export const defaultState = {
  bookmarks: new Map(),
  links: {}
};

export default function(
  state: IndexState = defaultState,
  action: ReduxActionWithPayload
): IndexState {
  switch (action.type) {
    case IndexActions.UPDATE: {
      const payload = action.payload as CollectionResponse;
      const collection: Index = parseCollection(payload.collection);

      return { ...state, links: collection.links };
    }

    case BookmarkActions.UPSERT: {
      const payload = action.payload as BookmarkDetails;
      const item: Bookmark = parseItem(payload.info.collection.items[0]);
      const history: BookmarkHistory = parseCollection(
        payload.history.collection
      );

      const bookmarks = new Map(state.bookmarks);
      let bookmark = bookmarks.get(item.id);

      bookmarks.set(item.id, {
        ...bookmark,
        ...item,
        history
      });

      bookmark = bookmarks.get(item.id);
      if (bookmark) {
        db.upsert(bookmark);
      }

      return { ...state, bookmarks };
    }

    case BookmarkActions.HISTORY: {
      const payload = action.payload as BookmarkHistoryPayload;
      const history: BookmarkHistory = parseCollection(
        payload.history.collection
      );

      const bookmarks = new Map(state.bookmarks);
      let bookmark = bookmarks.get(payload.bookmark.id);

      if (bookmark) {
        bookmarks.set(bookmark.id, {
          ...bookmark,
          history
        });

        bookmark = bookmarks.get(bookmark.id);
        if (bookmark) {
          db.upsert(bookmark);
        }
      }

      return { ...state, bookmarks };
    }

    /** intercept Feed actions to update the global bookmarks store */
    case FeedActions.APPEND:
    case FeedActions.PREPEND:
    case FeedActions.UPSERT_ALL:
    case FeedActions.REPLACE_ALL: {
      const payload = action.payload as CollectionResponse;
      const collection: Bookmarks = parseCollection(payload.collection);

      const bookmarks = new Map(state.bookmarks);

      collection.items.forEach(item => {
        let bookmark = bookmarks.get(item.id);
        bookmarks.set(item.id, {
          ...bookmark,
          ...item
        });
      });

      db.batch(Array.from(bookmarks.values()));

      return { ...state, bookmarks };
    }
  }

  return state;
}
