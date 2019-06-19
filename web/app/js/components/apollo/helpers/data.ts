import {
  FeedQueryData,
  FeedEventData,
  FeedEvent,
  FeedResults
} from "../../../types/feed";

export function getDataKey(data: FeedQueryData | FeedEventData): string | null {
  const keys = Object.keys(data);
  return keys.length > 0 ? keys[0] : null;
}

export function hasReceivedData(
  data: FeedQueryData | undefined
): [boolean, FeedResults] {
  let hasResults = false;
  let results: FeedResults = {
    first: "",
    last: "",
    results: [],
    total: 0,
    limit: 0
  };

  if (data) {
    const key = getDataKey(data);
    if (key && "results" in data[key]) {
      results = data[key];
      if (data[key].results.length > 0) {
        hasResults = true;
      }
    }
  }

  return [hasResults, results];
}

export function hasReceivedEvent(
  data: FeedEventData | undefined
): [boolean, FeedEvent | null] {
  let isReceived = false;
  let event: FeedEvent | null = null;

  if (data) {
    const key = getDataKey(data);
    if (key) {
      event = data[key];
      if (Object.keys(event).length > 0) {
        isReceived = true;
      }
    }
  }

  return [isReceived, event];
}
