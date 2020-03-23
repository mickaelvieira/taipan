import { Bookmark } from "../types/bookmark";
import { Subscription } from "../types/subscription";
import { Source } from "../types/syndication";

export function getBookmark(bookmark?: Partial<Bookmark>): Bookmark {
  return {
    id: "foo",
    url: new URL("https://foo.bar"),
    title: "bar",
    description: "baz",
    image: null,
    addedAt: new Date(),
    favoritedAt: new Date(),
    updatedAt: new Date(),
    isFavorite: false,
    ...bookmark,
  };
}

export function getSubscription(
  subscription?: Partial<Subscription>
): Subscription {
  return {
    id: "foo",
    url: new URL("https://foo.bar/baz"),
    domain: new URL("https://foo.bar"),
    title: "bar",
    type: "baz",
    frequency: "hourly",
    isSubscribed: true,
    createdAt: new Date(),
    updatedAt: new Date(),
    ...subscription,
  };
}

export function getSource(source?: Partial<Source>): Source {
  return {
    id: "foo",
    url: new URL("https://foo.bar/baz"),
    domain: new URL("https://foo.bar"),
    title: "baz",
    type: "baz",
    frequency: "hourly",
    isPaused: false,
    isDeleted: false,
    createdAt: new Date(),
    updatedAt: new Date(),
    parsedAt: new Date(),
    ...source,
  };
}
