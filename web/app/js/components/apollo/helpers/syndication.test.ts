import { addSource, removeSource } from "./syndication";
import { SyndicationResults, Source } from "../../../types/syndication";

function getSource(id: string): Source {
  return {
    id,
    url: "baz",
    domain: "foo",
    title: "baz",
    type: "baz",
    frequency: "hourly",
    isPaused: false,
    isDeleted: false,
    createdAt: "baz",
    updatedAt: "baz",
    parsedAt: "baz"
  };
}

describe("Syndication helpers", () => {
  describe("addSource", () => {
    let item: Source;
    let results: SyndicationResults;
    beforeEach(() => {
      results = {
        total: 0,
        offset: 3,
        limit: 10,
        results: []
      };
    });

    it("does not add falsy values", () => {
      results = addSource(results, item);
      expect(results.results.length).toEqual(0);
    });

    it("adds a item to a results", () => {
      const item1 = getSource("baz");
      results = addSource(results, item1);
      expect(results.results.length).toEqual(1);
    });

    it("adds items on top of the results", () => {
      const item1 = getSource("foo");
      const item2 = getSource("bar");
      const item3 = getSource("baz");
      results = addSource(results, item1);
      results = addSource(results, item2);
      results = addSource(results, item3);
      expect(results.results[0]).toEqual(item3);
      expect(results.results[1]).toEqual(item2);
      expect(results.results[2]).toEqual(item1);
    });

    it("clones existing items", () => {
      const item1 = getSource("foo");
      const item2 = getSource("bar");
      const item3 = getSource("baz");
      results = addSource(results, item1);
      results = addSource(results, item2);
      results = addSource(results, item3);
      expect(results.results[0]).toBe(item3);
      expect(results.results[1]).toEqual(item2);
      expect(results.results[1]).not.toBe(item2);
      expect(results.results[2]).toEqual(item1);
      expect(results.results[2]).not.toBe(item1);
    });

    it("updates result's total", () => {
      const item1 = getSource("baz");
      const item2 = getSource("bar");
      results = addSource(results, item1);
      results = addSource(results, item2);
      expect(results.total).toEqual(2);
    });

    it("does not duplicate items", () => {
      const item1 = getSource("baz");
      const item2 = getSource("baz");
      results = addSource(results, item1);
      results = addSource(results, item2);
      expect(results.total).toEqual(1);
      expect(results.results.length).toEqual(1);
    });
  });

  describe("removeSource", () => {
    const item1 = getSource("foo");
    const item2 = getSource("bar");
    const item3 = getSource("baz");
    let results: SyndicationResults;
    beforeEach(() => {
      results = {
        total: 3,
        offset: 3,
        limit: 10,
        results: [item1, item2, item3]
      };
    });

    it("removes a source from the results", () => {
      const item1 = getSource("baz");
      results = removeSource(results, item1);
      expect(results.results.length).toEqual(2);
    });

    it("clones existing sources", () => {
      const item1 = getSource("foo");
      results = removeSource(results, item1);
      expect(results.results[0]).toEqual(item2);
      expect(results.results[0]).not.toBe(item2);
      expect(results.results[1]).toEqual(item3);
      expect(results.results[1]).not.toBe(item3);
    });

    it("updates result's total", () => {
      const item1 = getSource("baz");
      const item2 = getSource("bar");
      results = removeSource(results, item1);
      results = removeSource(results, item2);
      expect(results.total).toEqual(1);
    });

    it("does not alter the results if the source is not present", () => {
      const item1 = getSource("foobar");
      results = removeSource(results, item1);
      expect(results).toBe(results);
    });
  });
});
