import { addSubscription, removeSubscription } from "./subscriptions";
import { SubscriptionResults, Subscription } from "../../../types/subscription";

function getSubscription(id: string): Subscription {
  return {
    id,
    url: "baz",
    title: "baz",
    type: "baz",
    isSubscribed: true,
    createdAt: "baz",
    updatedAt: "baz"
  };
}

describe("Syndication helpers", () => {
  describe("addSubscription", () => {
    let item: Subscription;
    let results: SubscriptionResults;
    beforeEach(() => {
      results = {
        total: 0,
        offset: 3,
        limit: 10,
        results: []
      };
    });

    it("does not add falsy values", () => {
      results = addSubscription(results, item);
      expect(results.results.length).toEqual(0);
    });

    it("adds a item to a results", () => {
      const item1 = getSubscription("baz");
      results = addSubscription(results, item1);
      expect(results.results.length).toEqual(1);
    });

    it("adds items on top of the results", () => {
      const item1 = getSubscription("foo");
      const item2 = getSubscription("bar");
      const item3 = getSubscription("baz");
      results = addSubscription(results, item1);
      results = addSubscription(results, item2);
      results = addSubscription(results, item3);
      expect(results.results[0]).toEqual(item3);
      expect(results.results[1]).toEqual(item2);
      expect(results.results[2]).toEqual(item1);
    });

    it("clones existing items", () => {
      const item1 = getSubscription("foo");
      const item2 = getSubscription("bar");
      const item3 = getSubscription("baz");
      results = addSubscription(results, item1);
      results = addSubscription(results, item2);
      results = addSubscription(results, item3);
      expect(results.results[0]).toBe(item3);
      expect(results.results[1]).toEqual(item2);
      expect(results.results[1]).not.toBe(item2);
      expect(results.results[2]).toEqual(item1);
      expect(results.results[2]).not.toBe(item1);
    });

    it("updates result's total", () => {
      const item1 = getSubscription("baz");
      const item2 = getSubscription("bar");
      results = addSubscription(results, item1);
      results = addSubscription(results, item2);
      expect(results.total).toEqual(2);
    });

    it("does not duplicate items", () => {
      const item1 = getSubscription("baz");
      const item2 = getSubscription("baz");
      results = addSubscription(results, item1);
      results = addSubscription(results, item2);
      expect(results.total).toEqual(1);
      expect(results.results.length).toEqual(1);
    });
  });

  describe("removeSubscription", () => {
    const item1 = getSubscription("foo");
    const item2 = getSubscription("bar");
    const item3 = getSubscription("baz");
    let results: SubscriptionResults;
    beforeEach(() => {
      results = {
        total: 3,
        offset: 3,
        limit: 10,
        results: [item1, item2, item3]
      };
    });

    it("removes a source from the results", () => {
      const item1 = getSubscription("baz");
      results = removeSubscription(results, item1);
      expect(results.results.length).toEqual(2);
    });

    it("clones existing sources", () => {
      const item1 = getSubscription("foo");
      results = removeSubscription(results, item1);
      expect(results.results[0]).toEqual(item2);
      expect(results.results[0]).not.toBe(item2);
      expect(results.results[1]).toEqual(item3);
      expect(results.results[1]).not.toBe(item3);
    });

    it("updates result's total", () => {
      const item1 = getSubscription("baz");
      const item2 = getSubscription("bar");
      results = removeSubscription(results, item1);
      results = removeSubscription(results, item2);
      expect(results.total).toEqual(1);
    });

    it("does not alter the results if the source is not present", () => {
      const item1 = getSubscription("foobar");
      results = removeSubscription(results, item1);
      expect(results).toBe(results);
    });
  });
});
