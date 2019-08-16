import { sortBy } from "lodash";
import { cloneDeep } from "lodash";
import {
  FeedQueryData,
  FeedItem,
  FeedName,
  FeedResults
} from "../../../types/feed";
import { ApolloClient } from "apollo-client";
import { getBoundaries, getDataKey, addItem, removeItemWithId } from "./feed";
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
      const results = sortBy(result.results as Bookmark[], ["addedAt"]);
      results.reverse();
      return {
        ...result,
        results
      };
    },
    favorites: (result: FeedResults): FeedResults => {
      const results = sortBy(result.results as Bookmark[], ["favoritedAt"]);
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

  appendLatest(documents: Document[]): void {
    const query = this.query["news"];

    try {
      const data = this.client.readQuery({ query }) as FeedQueryData;
      if (data) {
        const key = getDataKey(data);

        if (key) {
          const cloned = cloneDeep(data.feeds[key]);
          const total = cloned.total + documents.length;
          const results = cloned.results;

          // Append documents at the top of the feed
          documents.reverse();
          results.unshift(...documents);
          const [first, last] = getBoundaries(results);

          this.client.writeQuery({
            query,
            data: {
              feeds: {
                ...data.feeds,
                [key]: {
                  ...cloned,
                  first,
                  last,
                  total,
                  results
                }
              }
            }
          });
        }
      }
    } catch (e) {
      console.warn(e);
    }
  }
}
