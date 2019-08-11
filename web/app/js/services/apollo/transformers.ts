import { Document } from "../../types/document";
import { Bookmark } from "../../types/bookmark";
import { FetchResult } from "apollo-link";

export interface DocumentWithTypeName extends Document {
  __typename: string;
}

export interface BookmarkWithTypeName extends Bookmark {
  __typename: string;
}

type FeedItemWithTypename = DocumentWithTypeName | BookmarkWithTypeName;

export function transformerBookmark(
  bookmark: BookmarkWithTypeName
): BookmarkWithTypeName {
  if (typeof bookmark.url === "string") {
    bookmark.url = new URL(bookmark.url);
  }
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
}

export function transformDocument(
  document: DocumentWithTypeName
): DocumentWithTypeName {
  if (typeof document.url === "string") {
    document.url = new URL(document.url);
  }
  if (typeof document.createdAt === "string") {
    document.createdAt = new Date(document.createdAt);
  }
  if (typeof document.updatedAt === "string") {
    document.updatedAt = new Date(document.updatedAt);
  }
  return document;
}

function transform(items: FeedItemWithTypename[]): FeedItemWithTypename[] {
  return items.map(item => {
    if (item.__typename === "Document") {
      return transformDocument(item as DocumentWithTypeName);
    } else if (item.__typename === "Bookmark") {
      return transformerBookmark(item as BookmarkWithTypeName);
    } else {
      throw new Error(`Unknown type ${item.__typename}`);
    }
  });
}

export function transformFeedData(
  key: string,
  fetchResults: FetchResult
): FetchResult {
  fetchResults.data.feeds[key].results = transform(fetchResults.data.feeds[key]
    .results as FeedItemWithTypename[]);
  return fetchResults;
}

export function transformItemData(fetchResults: FetchResult): FetchResult {
  const key = Object.keys(fetchResults.data.bookmarks)[0];
  let item = fetchResults.data.bookmarks[key];
  if (item.__typename === "Document") {
    item = transformDocument(item as DocumentWithTypeName);
  } else if (item.__typename === "Bookmark") {
    item = transformerBookmark(item as BookmarkWithTypeName);
  } else {
    throw new Error(`Unknown type ${item.__typename}`);
  }
  fetchResults.data.bookmarks[key] = item;

  return fetchResults;
}
