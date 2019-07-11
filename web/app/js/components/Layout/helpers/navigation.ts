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
  "/account": "My Account"
};

export function getSectionTitle(path: string): string {
  return titles[path] ? titles[path] : "";
}

export function getPageTitle(path: string): string {
  return titles[path]
    ? `${appTitle} - ${appBaseline} - ${getSectionTitle(path)}`
    : "";
}
