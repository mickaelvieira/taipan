import { getDataKey } from "./feed";

describe("Feed helpers", () => {
  describe("getDataKey", () => {
    it("returns null when there is no feed data", () => {
      expect(getDataKey(undefined)).toBe(null);
    });

    test("retrieves the feed data key", () => {
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
});
