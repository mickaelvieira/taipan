import { getSectionTitle, getPageTitle } from "./navigation";

describe("Navigation helpers", () => {
  describe("getSectionTitle", () => {
    it("returns an empty string with falsy values", () => {
      expect(getSectionTitle("")).toBe("");
    });

    it("returns the news page title", () => {
      expect(getSectionTitle("/")).toBe("News");
    });

    it("returns the reading list page title", () => {
      expect(getSectionTitle("/reading-list")).toBe("Reading List");
    });

    it("returns the favorites page title", () => {
      expect(getSectionTitle("/favorites")).toBe("Favorites");
    });

    it("returns the syndication page title", () => {
      expect(getSectionTitle("/syndication")).toBe("RSS");
    });

    it("returns the account page title", () => {
      expect(getSectionTitle("/account")).toBe("My Account");
    });
  });

  describe("getPageTitle", () => {
    const appTitle = process.env.APP_NAME;
    const appBaseline = process.env.APP_BASELINE;

    it("returns an empty string with falsy values", () => {
      expect(getPageTitle("")).toBe("");
    });

    it("returns the news page title", () => {
      expect(getPageTitle("/")).toBe(`${appTitle} - ${appBaseline} - News`);
    });

    it("returns the reading list page title", () => {
      expect(getPageTitle("/reading-list")).toBe(
        `${appTitle} - ${appBaseline} - Reading List`
      );
    });

    it("returns the favorites page title", () => {
      expect(getPageTitle("/favorites")).toBe(
        `${appTitle} - ${appBaseline} - Favorites`
      );
    });

    it("returns the syndication page title", () => {
      expect(getPageTitle("/syndication")).toBe(
        `${appTitle} - ${appBaseline} - RSS`
      );
    });

    it("returns the account page title", () => {
      expect(getPageTitle("/account")).toBe(
        `${appTitle} - ${appBaseline} - My Account`
      );
    });
  });
});
