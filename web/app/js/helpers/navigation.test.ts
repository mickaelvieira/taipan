import { getSectionTitle, getPageTitle } from "./navigation";

describe("Navigation helpers", () => {
  describe("getSectionTitle", () => {
    it("returns an empty string with falsy values", () => {
      expect(getSectionTitle("")).toBe("");
    });

    it("returns the news section title", () => {
      expect(getSectionTitle("/")).toBe("News");
    });

    it("returns the reading list section title", () => {
      expect(getSectionTitle("/reading-list")).toBe("Reading List");
    });

    it("returns the favorites section title", () => {
      expect(getSectionTitle("/favorites")).toBe("Favorites");
    });

    it("returns the syndication section title", () => {
      expect(getSectionTitle("/syndication")).toBe("RSS");
    });

    it("returns the account section title", () => {
      expect(getSectionTitle("/account")).toBe("My Account");
    });

    it("returns the search results section title", () => {
      expect(getSectionTitle("/search")).toBe("Search");
    });
  });

  describe("getPageTitle", () => {
    const appTitle = process.env.APP_NAME;
    const appBaseline = process.env.APP_BASELINE;

    it("returns an empty string with falsy values", () => {
      expect(getPageTitle("")).toBe("");
    });

    it("returns the news page title", () => {
      expect(getPageTitle("/")).toBe(`News | ${appTitle} - ${appBaseline}`);
    });

    it("returns the reading list page title", () => {
      expect(getPageTitle("/reading-list")).toBe(
        `Reading List | ${appTitle} - ${appBaseline}`
      );
    });

    it("returns the favorites page title", () => {
      expect(getPageTitle("/favorites")).toBe(
        `Favorites | ${appTitle} - ${appBaseline}`
      );
    });

    it("returns the syndication page title", () => {
      expect(getPageTitle("/syndication")).toBe(
        `RSS | ${appTitle} - ${appBaseline}`
      );
    });

    it("returns the account page title", () => {
      expect(getPageTitle("/account")).toBe(
        `My Account | ${appTitle} - ${appBaseline}`
      );
    });

    it("returns the search results page title", () => {
      expect(getPageTitle("/search")).toBe(
        `Search | ${appTitle} - ${appBaseline}`
      );
    });
  });
});
