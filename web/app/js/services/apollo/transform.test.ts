import { FetchResult } from "apollo-link";
import transformer from "./transform";
import { Document } from "../../types/document";
import { Bookmark } from "../../types/bookmark";
import { Subscription } from "../../types/subscription";
import { Source, Tag } from "../../types/syndication";

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

function getFetchedTag(o?: Record<string, any>): any {
  return {
    id: "foo",
    label: "bar",
    createdAt: "2019-07-28T11:15:39Z",
    updatedAt: "2019-07-28T11:15:39Z",
    __typename: "SyndicationTag",
    ...o
  };
}

function getFetchEmail(o?: Record<string, any>): any {
  return {
    id: "foo",
    value: "foo@bar.baz",
    isPrimary: false,
    isConfirmed: false,
    createdAt: "2019-07-28T11:15:39Z",
    updatedAt: "2019-07-28T11:15:39Z",
    __typename: "Email",
    ...o
  };
}

function getFetchUser(o?: Record<string, any>): any {
  return {
    id: "foo",
    firstname: "foo",
    lastname: "bar",
    emails: [getFetchEmail()],
    createdAt: "2019-07-28T11:15:39Z",
    updatedAt: "2019-07-28T11:15:39Z",
    __typename: "User",
    ...o
  };
}

describe("Transformer", () => {
  let result: FetchResult;
  beforeAll(() => {
    result = {
      data: {
        foonull: null,
        barnull: {
          bar: null
        },
        document: {
          bar: getFetchedDocument()
        },
        image: {
          bar: getFetchedDocument({
            image: { url: "https://foo.bar/baz.jpeg", __typename: "Image" }
          })
        },
        user: {
          bar: getFetchUser()
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
        tag: {
          bar: getFetchedTag()
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
        sources: {
          foo: {
            results: [
              getFetchedSource({ id: "foo" }),
              getFetchedSource({ id: "bar" }),
              getFetchedSource({ id: "baz" })
            ],
            __typename: "SourceCollection"
          }
        },
        tags: {
          foo: {
            results: [
              getFetchedTag({ id: "foo" }),
              getFetchedTag({ id: "bar" }),
              getFetchedTag({ id: "baz" })
            ],
            __typename: "SyndicationTagCollection"
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

  it("handles null data", () => {
    const data = transformer(result);
    if (!data.data) {
      throw new Error("No data returned");
    }
    expect(data.data.foonull).toBe(null);
    expect(data.data.barnull.bar).toBe(null);
  });

  it("does not transform unknown types", () => {
    const data = transformer(result);
    if (!data.data) {
      throw new Error("No data returned");
    }
    const unknown = data.data.unknown.foo;
    expect(unknown.bar).toBe("bar");
    expect(unknown.baz).toBe("baz");
  });

  it("transform a single document", () => {
    const data = transformer(result);
    if (!data.data) {
      throw new Error("No data returned");
    }
    const document = data.data.document.bar;
    expect(document.url instanceof URL).toBe(true);
    expect(document.createdAt instanceof Date).toBe(true);
    expect(document.updatedAt instanceof Date).toBe(true);
  });

  it("transform a nested types", () => {
    const data = transformer(result);
    if (!data.data) {
      throw new Error("No data returned");
    }
    const document = data.data.image.bar;
    expect(document.image.url instanceof URL).toBe(true);
  });

  it("transform a single user", () => {
    const data = transformer(result);
    if (!data.data) {
      throw new Error("No data returned");
    }
    const user = data.data.user.bar;
    expect(user.createdAt instanceof Date).toBe(true);
    expect(user.updatedAt instanceof Date).toBe(true);
    expect(user.emails[0].createdAt instanceof Date).toBe(true);
    expect(user.emails[0].updatedAt instanceof Date).toBe(true);
  });

  it("transform a single bookmark", () => {
    const data = transformer(result);
    if (!data.data) {
      throw new Error("No data returned");
    }
    const bookmark = data.data.bookmark.bar;
    expect(bookmark.url instanceof URL).toBe(true);
    expect(bookmark.addedAt instanceof Date).toBe(true);
    expect(bookmark.favoritedAt instanceof Date).toBe(true);
    expect(bookmark.updatedAt instanceof Date).toBe(true);
  });

  it("transform a single source", () => {
    const data = transformer(result);
    if (!data.data) {
      throw new Error("No data returned");
    }
    const source = data.data.source.bar;
    expect(source.url instanceof URL).toBe(true);
    expect(source.domain instanceof URL).toBe(true);
    expect(source.parsedAt instanceof Date).toBe(true);
    expect(source.createdAt instanceof Date).toBe(true);
    expect(source.updatedAt instanceof Date).toBe(true);
  });

  it("transform a single subscription", () => {
    const data = transformer(result);
    if (!data.data) {
      throw new Error("No data returned");
    }
    const subscription = data.data.subscription.bar;
    expect(subscription.url instanceof URL).toBe(true);
    expect(subscription.domain instanceof URL).toBe(true);
    expect(subscription.createdAt instanceof Date).toBe(true);
    expect(subscription.updatedAt instanceof Date).toBe(true);
  });

  it("transform a single tag", () => {
    const data = transformer(result);
    if (!data.data) {
      throw new Error("No data returned");
    }
    const tag = data.data.tag.bar;
    expect(tag.createdAt instanceof Date).toBe(true);
    expect(tag.updatedAt instanceof Date).toBe(true);
  });

  it("transform a collection of documents", () => {
    const data = transformer(result);
    if (!data.data) {
      throw new Error("No data returned");
    }
    const documents = data.data.feeddocuments.documents.results;
    documents.forEach((document: Document) => {
      expect(document.url instanceof URL).toBe(true);
      expect(document.createdAt instanceof Date).toBe(true);
      expect(document.updatedAt instanceof Date).toBe(true);
    });
  });

  it("transform a search result of documents", () => {
    const data = transformer(result);
    if (!data.data) {
      throw new Error("No data returned");
    }
    const documents = data.data.searchdocuments.documents.results;
    documents.forEach((document: Document) => {
      expect(document.url instanceof URL).toBe(true);
      expect(document.createdAt instanceof Date).toBe(true);
      expect(document.updatedAt instanceof Date).toBe(true);
    });
  });

  it("transform a collection of bookmarks", () => {
    const data = transformer(result);
    if (!data.data) {
      throw new Error("No data returned");
    }
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
    if (!data.data) {
      throw new Error("No data returned");
    }
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
    if (!data.data) {
      throw new Error("No data returned");
    }
    const subscriptions = data.data.subscriptions.foo.results;
    subscriptions.forEach((subscription: Subscription) => {
      expect(subscription.url instanceof URL).toBe(true);
      expect(subscription.domain instanceof URL).toBe(true);
      expect(subscription.createdAt instanceof Date).toBe(true);
      expect(subscription.updatedAt instanceof Date).toBe(true);
    });
  });

  it("transform a collection of syndication sources", () => {
    const data = transformer(result);
    if (!data.data) {
      throw new Error("No data returned");
    }
    const sources = data.data.sources.foo.results;
    sources.forEach((source: Source) => {
      expect(source.url instanceof URL).toBe(true);
      expect(source.domain instanceof URL).toBe(true);
      expect(source.createdAt instanceof Date).toBe(true);
      expect(source.updatedAt instanceof Date).toBe(true);
      expect(source.parsedAt instanceof Date).toBe(true);
    });
  });

  it("transform a collection of syndication tags", () => {
    const data = transformer(result);
    if (!data.data) {
      throw new Error("No data returned");
    }
    const tags = data.data.tags.foo.results;
    tags.forEach((tag: Tag) => {
      expect(tag.createdAt instanceof Date).toBe(true);
      expect(tag.updatedAt instanceof Date).toBe(true);
    });
  });
});
