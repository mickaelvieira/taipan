import {
  getDataKey,
  hasReceivedData,
  hasReceivedEvent,
  feedResultsAction,
  getBoundaries
} from "./feed";
import { Bookmark } from "../../../types/bookmark";
import {
  FeedResults,
  FeedItem,
  FeedEvent,
  FeedEventData
} from "../../../types/feed";

const addItem = feedResultsAction["Add"];
const removeItem = feedResultsAction["Remove"];

function getBookmark(id: string): Bookmark {
  return {
    id,
    url: "baz",
    title: "baz",
    description: "baz",
    image: null,
    addedAt: "baz",
    updatedAt: "baz",
    isFavorite: false
  };
}

describe("Feed helpers", () => {
  describe("getDataKey", () => {
    it("returns null when there is no feed data", () => {
      expect(getDataKey(undefined)).toBe(null);
    });

    it("retrieves the feed data key", () => {
      expect(
        getDataKey({
          feeds: {
            foo: {
              total: 0,
              first: "",
              last: "",
              limit: 0,
              results: []
            }
          }
        })
      ).toBe("foo");
    });
  });

  describe("hasReceivedData", () => {
    it("returns an empty result by default", () => {
      const [hasResults, results] = hasReceivedData(undefined);
      expect(hasResults).toBe(false);
      expect(results).toEqual({
        total: 0,
        first: "",
        last: "",
        limit: 0,
        results: []
      });
    });

    it("has no results when the result has no data", () => {
      const result = {
        feeds: {
          foo: {
            total: 0,
            first: "foo",
            last: "bar",
            limit: 10,
            results: []
          }
        }
      };
      const [hasResults, results] = hasReceivedData(result);
      expect(hasResults).toBe(false);
      expect(results).toEqual({
        total: 0,
        first: "foo",
        last: "bar",
        limit: 10,
        results: []
      });
    });

    it("has results when the result has data", () => {
      const item = getBookmark("baz");
      const result = {
        feeds: {
          foo: {
            total: 0,
            first: "foo",
            last: "bar",
            limit: 10,
            results: [item]
          }
        }
      };
      const [hasResults, results] = hasReceivedData(result);
      expect(hasResults).toBe(true);
      expect(results).toEqual({
        total: 0,
        first: "foo",
        last: "bar",
        limit: 10,
        results: [item]
      });
    });
  });

  describe("hasReceivedEvent", () => {
    let event: FeedEvent;
    let result: FeedEventData;
    beforeAll(() => {
      event = {
        id: "foo",
        emitter: "baz",
        item: getBookmark("baz"),
        action: "Add",
        topic: "Favorites"
      };
      result = {
        foo: event
      };
    });

    it("returns null by default", () => {
      const [hasEvent, e] = hasReceivedEvent(undefined);
      expect(hasEvent).toBe(false);
      expect(e).toBe(null);
    });

    it("returns the event", () => {
      const [hasEvent, e] = hasReceivedEvent(result);
      expect(hasEvent).toBe(true);
      expect(e).toEqual(event);
    });
  });

  describe("getBoundaries", () => {
    it("returns empty boundaries when there is no results", () => {
      const [first, last] = getBoundaries([]);
      expect(first).toBe("");
      expect(last).toBe("");
    });

    it("returns the first and the last ID", () => {
      const item1 = getBookmark("foo");
      const item2 = getBookmark("bar");
      const item3 = getBookmark("baz");
      const [first, last] = getBoundaries([item1, item2, item3]);
      expect(first).toBe("foo");
      expect(last).toBe("baz");
    });
  });

  describe("addItem", () => {
    let feed: FeedResults;
    let item: FeedItem;
    beforeEach(() => {
      feed = {
        total: 0,
        first: "",
        last: "",
        limit: 10,
        results: []
      };
    });

    it("does not add falsy values", () => {
      feed = addItem(feed, item);
      expect(feed.results.length).toEqual(0);
    });

    it("adds a item to a feed", () => {
      const item1 = getBookmark("baz");
      feed = addItem(feed, item1);
      expect(feed.results.length).toEqual(1);
    });

    it("adds items on top of the feed", () => {
      const item1 = getBookmark("foo");
      const item2 = getBookmark("bar");
      const item3 = getBookmark("baz");
      feed = addItem(feed, item1);
      feed = addItem(feed, item2);
      feed = addItem(feed, item3);
      expect(feed.results[0]).toEqual(item3);
      expect(feed.results[1]).toEqual(item2);
      expect(feed.results[2]).toEqual(item1);
    });

    it("updates result's boundaries", () => {
      const item1 = getBookmark("foo");
      const item2 = getBookmark("bar");
      const item3 = getBookmark("baz");
      feed = addItem(feed, item1);
      feed = addItem(feed, item2);
      feed = addItem(feed, item3);
      expect(feed.first).toEqual("baz");
      expect(feed.last).toEqual("foo");
    });

    it("clones existing items", () => {
      const item1 = getBookmark("foo");
      const item2 = getBookmark("bar");
      const item3 = getBookmark("baz");
      feed = addItem(feed, item1);
      feed = addItem(feed, item2);
      feed = addItem(feed, item3);
      expect(feed.results[0]).toBe(item3);
      expect(feed.results[1]).toEqual(item2);
      expect(feed.results[1]).not.toBe(item2);
      expect(feed.results[2]).toEqual(item1);
      expect(feed.results[2]).not.toBe(item1);
    });

    it("updates result's total", () => {
      const item1 = getBookmark("baz");
      const item2 = getBookmark("bar");
      feed = addItem(feed, item1);
      feed = addItem(feed, item2);
      expect(feed.total).toEqual(2);
    });

    it("does not duplicate items", () => {
      const item1 = getBookmark("baz");
      const item2 = getBookmark("baz");
      feed = addItem(feed, item1);
      feed = addItem(feed, item2);
      expect(feed.total).toEqual(1);
      expect(feed.results.length).toEqual(1);
    });
  });

  describe("removeItem", () => {
    const item1 = getBookmark("foo");
    const item2 = getBookmark("bar");
    const item3 = getBookmark("baz");
    let feed: FeedResults;
    beforeEach(() => {
      feed = {
        total: 3,
        first: "foo",
        last: "baz",
        limit: 10,
        results: [item1, item2, item3]
      };
    });

    it("removes an item from the feed", () => {
      const item1 = getBookmark("baz");
      feed = removeItem(feed, item1);
      expect(feed.results.length).toEqual(2);
    });

    it("updates result's upper boundary", () => {
      const item1 = getBookmark("foo");
      feed = removeItem(feed, item1);
      expect(feed.first).toEqual("bar");
    });

    it("updates result's lower boundary", () => {
      const item1 = getBookmark("baz");
      feed = removeItem(feed, item1);
      expect(feed.last).toEqual("bar");
    });

    it("clones existing items", () => {
      const item1 = getBookmark("foo");
      feed = removeItem(feed, item1);
      expect(feed.results[0]).toEqual(item2);
      expect(feed.results[0]).not.toBe(item2);
      expect(feed.results[1]).toEqual(item3);
      expect(feed.results[1]).not.toBe(item3);
    });

    it("updates result's total", () => {
      const item1 = getBookmark("baz");
      const item2 = getBookmark("bar");
      feed = removeItem(feed, item1);
      feed = removeItem(feed, item2);
      expect(feed.total).toEqual(1);
    });

    it("does not alter the feed if the item is not present", () => {
      const item1 = getBookmark("foobar");
      feed = removeItem(feed, item1);
      expect(feed).toBe(feed);
    });
  });
});
