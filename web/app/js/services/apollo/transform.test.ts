import { FetchResult } from "apollo-link";
import transformer from "./transform";
import { Document } from "../../types/document";
import { Bookmark } from "../../types/bookmark";
import { Subscription } from "../../types/subscription";

/* eslint @typescript-eslint/no-explicit-any: "off" */

function getFetchedDocument(o?: Record<string, any>): any {
  return {
    id: "foo",
    url: "https://foo.bar/baz",
    image: null,
    createdAt: "2019-07-28T11:15:39Z",
    updatedAt: "2019-07-28T11:15:39Z",
    __typename: "Document",
    ...o
  };
}

function getFetchedBookmark(o?: Record<string, any>): any {
  return {
    id: "foo",
    url: "https://foo.bar",
    image: null,
    addedAt: "2019-07-28T11:15:39Z",
    favoritedAt: "2019-07-28T11:15:39Z",
    updatedAt: "2019-07-28T11:15:39Z",
    __typename: "Bookmark",
    ...o
  };
}

function getFetchedSource(o?: Record<string, any>): any {
  return {
    id: "foo",
    url: "https://foo.bar/baz",
    domain: "https://foo.bar",
    parsedAt: "2019-07-28T11:15:39Z",
    createdAt: "2019-07-28T11:15:39Z",
    updatedAt: "2019-07-28T11:15:39Z",
    __typename: "Source",
    ...o
  };
}

function getFetchedSubscription(o?: Record<string, any>): any {
  return {
    id: "foo",
    url: "https://foo.bar/baz",
    domain: "https://foo.bar",
    createdAt: "2019-07-28T11:15:39Z",
    updatedAt: "2019-07-28T11:15:39Z",
    __typename: "UserSubscription",
    ...o
  };
}

describe("Transformer", () => {
  let result: FetchResult;
  beforeAll(() => {
    result = {
      data: {
        document: {
          bar: getFetchedDocument()
        },
        image: {
          bar: getFetchedDocument({
            image: { url: "https://foo.bar/baz.jpeg", __typename: "Image" }
          })
        },
        bookmark: {
          bar: getFetchedBookmark()
        },
        source: {
          bar: getFetchedSource()
        },
        subscription: {
          bar: getFetchedSubscription()
        },
        feeddocuments: {
          documents: {
            results: [
              getFetchedDocument({ id: "foo" }),
              getFetchedDocument({ id: "bar" }),
              getFetchedDocument({ id: "baz" })
            ],
            __typename: "FeedDocumentResults"
          }
        },
        feedbookmarks: {
          bookmarks: {
            results: [
              getFetchedBookmark({ id: "foo" }),
              getFetchedBookmark({ id: "bar" }),
              getFetchedBookmark({ id: "baz" })
            ],
            __typename: "FeedBookmarkResults"
          }
        },
        searchdocuments: {
          documents: {
            results: [
              getFetchedDocument({ id: "foo" }),
              getFetchedDocument({ id: "bar" }),
              getFetchedDocument({ id: "baz" })
            ],
            __typename: "DocumentSearchResults"
          }
        },
        searchbookmarks: {
          bookmarks: {
            results: [
              getFetchedBookmark({ id: "foo" }),
              getFetchedBookmark({ id: "bar" }),
              getFetchedBookmark({ id: "baz" })
            ],
            __typename: "BookmarkSearchResults"
          }
        },
        subscriptions: {
          foo: {
            results: [
              getFetchedSubscription({ id: "foo" }),
              getFetchedSubscription({ id: "bar" }),
              getFetchedSubscription({ id: "baz" })
            ],
            __typename: "SubscriptionCollection"
          }
        },
        unknown: {
          foo: {
            bar: "bar",
            baz: "baz"
          }
        }
      }
    };
  });

  it("does not transform unknown types", () => {
    const data = transformer(result);
    const unknown = data.data.unknown.foo;
    expect(unknown.bar).toBe("bar");
    expect(unknown.baz).toBe("baz");
  });

  it("transform a single document", () => {
    const data = transformer(result);
    const document = data.data.document.bar;
    expect(document.url instanceof URL).toBe(true);
    expect(document.createdAt instanceof Date).toBe(true);
    expect(document.updatedAt instanceof Date).toBe(true);
  });

  it("transform a nested types", () => {
    const data = transformer(result);
    const document = data.data.image.bar;
    expect(document.image.url instanceof URL).toBe(true);
  });

  it("transform a single bookmark", () => {
    const data = transformer(result);
    const bookmark = data.data.bookmark.bar;
    expect(bookmark.url instanceof URL).toBe(true);
    expect(bookmark.addedAt instanceof Date).toBe(true);
    expect(bookmark.favoritedAt instanceof Date).toBe(true);
    expect(bookmark.updatedAt instanceof Date).toBe(true);
  });

  it("transform a single source", () => {
    const data = transformer(result);
    const source = data.data.source.bar;
    expect(source.url instanceof URL).toBe(true);
    expect(source.domain instanceof URL).toBe(true);
    expect(source.parsedAt instanceof Date).toBe(true);
    expect(source.createdAt instanceof Date).toBe(true);
    expect(source.updatedAt instanceof Date).toBe(true);
  });

  it("transform a single subscription", () => {
    const data = transformer(result);
    const subscription = data.data.subscription.bar;
    expect(subscription.url instanceof URL).toBe(true);
    expect(subscription.domain instanceof URL).toBe(true);
    expect(subscription.createdAt instanceof Date).toBe(true);
    expect(subscription.updatedAt instanceof Date).toBe(true);
  });

  it("transform a collection of documents", () => {
    const data = transformer(result);
    const documents = data.data.feeddocuments.documents.results;
    documents.forEach((document: Document) => {
      expect(document.url instanceof URL).toBe(true);
      expect(document.createdAt instanceof Date).toBe(true);
      expect(document.updatedAt instanceof Date).toBe(true);
    });
  });

  it("transform a search result of documents", () => {
    const data = transformer(result);
    const documents = data.data.searchdocuments.documents.results;
    documents.forEach((document: Document) => {
      expect(document.url instanceof URL).toBe(true);
      expect(document.createdAt instanceof Date).toBe(true);
      expect(document.updatedAt instanceof Date).toBe(true);
    });
  });

  it("transform a collection of bookmarks", () => {
    const data = transformer(result);
    const bookmarks = data.data.feedbookmarks.bookmarks.results;
    bookmarks.forEach((bookmark: Bookmark) => {
      expect(bookmark.url instanceof URL).toBe(true);
      expect(bookmark.addedAt instanceof Date).toBe(true);
      expect(bookmark.favoritedAt instanceof Date).toBe(true);
      expect(bookmark.updatedAt instanceof Date).toBe(true);
    });
  });

  it("transform a search result of bookmarks", () => {
    const data = transformer(result);
    const bookmarks = data.data.searchbookmarks.bookmarks.results;
    bookmarks.forEach((bookmark: Bookmark) => {
      expect(bookmark.url instanceof URL).toBe(true);
      expect(bookmark.addedAt instanceof Date).toBe(true);
      expect(bookmark.favoritedAt instanceof Date).toBe(true);
      expect(bookmark.updatedAt instanceof Date).toBe(true);
    });
  });

  it("transform a collection of subscriptions", () => {
    const data = transformer(result);
    const subscriptions = data.data.subscriptions.foo.results;
    subscriptions.forEach((subscription: Subscription) => {
      expect(subscription.url instanceof URL).toBe(true);
      expect(subscription.domain instanceof URL).toBe(true);
      expect(subscription.createdAt instanceof Date).toBe(true);
      expect(subscription.updatedAt instanceof Date).toBe(true);
    });
  });
});
