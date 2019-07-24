import { sortBy } from "lodash";
import {
  FeedQueryData,
  FeedItem,
  FeedName,
  FeedResults
} from "../../../types/feed";
import { ApolloClient } from "apollo-client";
import { getDataKey, addItem, removeItemWithId } from "./feed";
import { Bookmark } from "../../../types/bookmark";
import { Document } from "../../../types/document";
import { queryNews, queryReadingList, queryFavorites } from "../Query/Feed";

// @TODO workaround for now, since apollo client does not support custom scalar transformation
function transformBookmarks(bookmarks: Bookmark[]): Bookmark[] {
  return bookmarks.map(bookmark => {
    if (typeof bookmark.addedAt === "string") {
      bookmark.addedAt = new Date(bookmark.addedAt);
    }
    if (typeof bookmark.favoritedAt === "string") {
      bookmark.favoritedAt = new Date(bookmark.favoritedAt);
    }
    if (typeof bookmark.updatedAt === "string") {
      bookmark.updatedAt = new Date(bookmark.updatedAt);
    }
    return bookmark;
  });
}

// @TODO workaround for now, since apollo client does not support custom scalar transformation
// function transformDocuments(documents: Document[]): Document[] {
//   return documents.map(document => {
//     if (typeof document.createdAt === "string") {
//       document.createdAt = new Date(document.createdAt);
//     }
//     if (typeof document.updatedAt === "string") {
//       document.updatedAt = new Date(document.updatedAt);
//     }
//     return document;
//   });
// }

export default class FeedsUpdater {
  client: ApolloClient<object>;

  query = {
    news: queryNews,
    readinglist: queryReadingList,
    favorites: queryFavorites
  };

  sorting = {
    news: (result: FeedResults): FeedResults => {
      const results = sortBy(result.results, ["id"]);
      results.reverse();
      return {
        ...result,
        results
      };
    },
    readinglist: (result: FeedResults): FeedResults => {
      const results = sortBy(transformBookmarks(result.results as Bookmark[]), [
        "addedAt"
      ]);
      results.reverse();
      return {
        ...result,
        results
      };
    },
    favorites: (result: FeedResults): FeedResults => {
      const results = sortBy(transformBookmarks(result.results as Bookmark[]), [
        "favoritedAt"
      ]);
      results.reverse();
      return {
        ...result,
        results
      };
    }
  };

  constructor(client: ApolloClient<object>) {
    this.client = client;
  }

  private addTo(name: FeedName, item: FeedItem): void {
    const query = this.query[name];
    const sort = this.sorting[name];
    try {
      const data = this.client.readQuery({ query }) as FeedQueryData;
      if (data) {
        const key = getDataKey(data);
        if (key) {
          const result = sort(addItem(data.feeds[key], item));
          this.client.writeQuery({
            query,
            data: {
              feeds: {
                ...data.feeds,
                [key]: result
              }
            }
          });
        }
      }
    } catch (e) {
      // console.warn(e.message);
    }
  }

  private removeFrom(name: FeedName, id: string): void {
    const query = this.query[name];
    try {
      const data = this.client.readQuery({ query }) as FeedQueryData;
      if (data) {
        const key = getDataKey(data);
        if (key) {
          const result = removeItemWithId(data.feeds[key], id);
          this.client.writeQuery({
            query,
            data: {
              feeds: {
                ...data.feeds,
                [key]: result
              }
            }
          });
        }
      }
    } catch (e) {
      // console.warn(e.message);
    }
  }

  bookmark(bookmark: Bookmark): void {
    const name = bookmark.isFavorite ? "favorites" : "readinglist";
    this.addTo(name, bookmark);
    this.removeFrom("news", bookmark.id);
  }

  unbookmark(document: Document): void {
    this.removeFrom("favorites", document.id);
    this.removeFrom("readinglist", document.id);
    this.addTo("news", document);
  }

  favorite(bookmark: Bookmark): void {
    this.addTo("favorites", bookmark);
    this.removeFrom("readinglist", bookmark.id);
  }

  unfavorite(bookmark: Bookmark): void {
    this.addTo("readinglist", bookmark);
    this.removeFrom("favorites", bookmark.id);
  }
}
