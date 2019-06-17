import { FeedData, FeedItem } from "../../../types/feed";

export function getDataKey(data: FeedData): string | null {
  const keys = Object.keys(data);
  return keys.length > 0 ? keys[0] : null;
}

export function hasReceivedData(
  data: FeedData | undefined
): [boolean, FeedItem[]] {
  let hasResults = false;
  let results: FeedItem[] = [];

  if (data) {
    const key = getDataKey(data);
    if (key && "results" in data[key]) {
      results = data[key].results;
      if (results.length > 0) {
        hasResults = true;
      }
    }
  }

  return [hasResults, results];
}
