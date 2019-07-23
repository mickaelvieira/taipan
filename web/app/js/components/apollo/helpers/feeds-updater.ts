import { FeedQueryData, FeedItem, FeedName } from "../../../types/feed";
import { ApolloClient } from "apollo-client";
import { getDataKey, addItem, removeItemWithId } from "./feed";
import { Bookmark } from "../../../types/bookmark";
import { Document } from "../../../types/document";
import { queryNews, queryReadingList, queryFavorites } from "../Query/Feed";

export default class FeedsUpdater {
  client: ApolloClient<object>;

  query = {
    news: queryNews,
    readinglist: queryReadingList,
    favorites: queryFavorites
  };

  constructor(client: ApolloClient<object>) {
    this.client = client;
  }

  private addTo(name: FeedName, item: FeedItem): void {
    const query = this.query[name];
    try {
      const data = this.client.readQuery({ query }) as FeedQueryData;
      if (data) {
        const key = getDataKey(data);
        if (key) {
          const result = addItem(data.feeds[key], item);
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
      console.warn(e.message);
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
