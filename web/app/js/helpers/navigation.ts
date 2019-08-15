interface Titles {
  [path: string]: string;
}

const appTitle = process.env.APP_NAME;
const appBaseline = process.env.APP_BASELINE;

const titles: Titles = {
  "/": "News",
  "/reading-list": "Reading List",
  "/favorites": "Favorites",
  "/syndication": "RSS",
  "/account": "My Account",
  "/search": "Search"
};

export function getSectionTitle(pathname: string): string {
  return titles[pathname] ? titles[pathname] : "";
}

export function getPageTitle(pathname: string): string {
  return titles[pathname]
    ? `${appTitle} - ${appBaseline} - ${getSectionTitle(pathname)}`
    : "";
}

export function isSearchResultsPage(pathname: string): boolean {
  return pathname === "/search";
}

export class Page {
  url: URL;

  constructor(url: string) {
    this.url = new URL(url);
  }
  getTitle(): string {
    return getPageTitle(this.url.pathname);
  }
  getSection(): string {
    return getSectionTitle(this.url.pathname);
  }
  isSearchResults(): boolean {
    return isSearchResultsPage(this.url.pathname);
  }
}