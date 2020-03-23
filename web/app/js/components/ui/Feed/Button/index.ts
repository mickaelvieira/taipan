import FavoriteButton from "./Favorite";
import UnfavoriteButton from "./Unfavorite";
import BookmarkButton from "./Bookmark";
import RefreshButton from "./Refresh";
import UnbookmarkButton from "./Unbookmark";
import BookmarkAndFavoriteButton from "./BookmarkAndFavorite";
import ShareButton from "./Share";
import { Undoer, CacheUpdater } from "../../../../types";
import { FeedItem } from "../../../../types/feed";

export interface SuccessOptions {
  undo: Undoer;
  item: FeedItem;
  updateCache: CacheUpdater;
}

export {
  FavoriteButton,
  UnfavoriteButton,
  BookmarkButton,
  RefreshButton,
  UnbookmarkButton,
  BookmarkAndFavoriteButton,
  ShareButton,
};
