import { Bookmarks, BookmarksLinks } from "types/bookmark";
import { parseCollection, getAllIds } from "collection";
import { CollectionResponse } from "collection/types";
import {
  FeedActions,
  ReduxActionWithPayload,
  ParamId
} from "store/actions/types";
import feed, { FeedTypes } from "services/repository/feed";

export interface FeedState {
  items: string[];
  links: BookmarksLinks;
}

export const defaultState = {
  items: [],
  links: {}
};

export default function(
  state: FeedState = defaultState,
  action: ReduxActionWithPayload
): FeedState {
  switch (action.type) {
    case FeedActions.REPLACE_ALL: {
      const payload = action.payload as CollectionResponse;
      const collection: Bookmarks = parseCollection(payload.collection);
      const items = getAllIds(collection);

      feed.update({
        type: FeedTypes.LATEST,
        results: items
      });

      return {
        ...state,
        items,
        links: collection.links
      };
    }
    case FeedActions.APPEND: {
      const payload = action.payload as CollectionResponse;
      const collection: Bookmarks = parseCollection(payload.collection);
      const items = [...state.items];

      items.push(...getAllIds(collection));

      feed.update({
        type: FeedTypes.LATEST,
        results: items
      });

      return {
        ...state,
        items,
        links: collection.links
      };
    }
    case FeedActions.PREPEND: {
      const payload = action.payload as CollectionResponse;
      const collection: Bookmarks = parseCollection(payload.collection);
      const items = [...state.items];

      items.unshift(...getAllIds(collection));

      feed.update({
        type: FeedTypes.LATEST,
        results: items
      });

      return {
        ...state,
        items
      };
    }
    case FeedActions.DELETE: {
      const payload = action.payload as ParamId;
      const items = [...state.items];
      const index = items.indexOf(payload.id);

      if (index >= 0) {
        delete items[index];
      }

      feed.update({
        type: FeedTypes.LATEST,
        results: items
      });

      return {
        ...state,
        items
      };
    }
  }

  return state;
}
