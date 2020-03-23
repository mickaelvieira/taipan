import { SearchQueryData, SearchResults } from "../../../types/search";

export function getDataKey(data: SearchQueryData | undefined): string | null {
  if (typeof data === "undefined") {
    return null;
  }
  const keys = Object.keys(data);
  if (keys.length === 0) {
    return null;
  }
  return keys[0];
}

export function hasReceivedData(
  data: SearchQueryData | undefined
): [boolean, SearchResults] {
  let hasResults = false;
  let results: SearchResults = {
    results: [],
    total: 0,
    offset: 0,
    limit: 0,
  };

  console.log(data);
  if (data) {
    const key = getDataKey(data);
    if (key) {
      console.log(key);
      results = data[key].search;

      console.log(results);
      if (results.results.length > 0) {
        hasResults = true;
      }
    }
  }

  return [hasResults, results];
}
