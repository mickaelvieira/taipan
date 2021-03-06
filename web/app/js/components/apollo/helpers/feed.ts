import { cloneDeep } from "lodash";
import {
  FeedQueryData,
  FeedEventData,
  FeedResults,
  FeedItem,
  FeedEvent,
} from "../../../types/feed";

export function getDataKey(data: FeedQueryData | undefined): string | null {
  if (typeof data === "undefined") {
    return null;
  }
  const keys = "feeds" in data ? Object.keys(data.feeds) : [];
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
    limit: 0,
  };

  if (data) {
    const key = getDataKey(data);
    if (key) {
      results = data.feeds[key];
      if (results.results.length > 0) {
        hasResults = true;
      }
    }
  }

  return [hasResults, results];
}

function getEventKey(data: FeedEventData): string | null {
  const keys = Object.keys(data);
  return keys.length > 0 ? keys[0] : null;
}

export function hasReceivedEvent(
  data: FeedEventData | undefined
): [boolean, FeedEvent | null] {
  let isReceived = false;
  let event: FeedEvent | null = null;

  if (data) {
    const key = getEventKey(data);
    if (key) {
      event = data[key];
      if (Object.keys(event).length > 0) {
        isReceived = true;
      }
    }
  }

  return [isReceived, event];
}

export function getBoundaries(results: FeedItem[]): [string, string] {
  let first = "";
  let last = "";
  if (results.length > 0) {
    first = results[0].id;
    last = results[results.length - 1].id;
  }
  return [first, last];
}

export function hasItem(result: FeedResults, item: FeedItem): boolean {
  const index = result.results.findIndex(({ id }) => id === item.id);
  return index >= 0;
}

export function addItem(result: FeedResults, item: FeedItem): FeedResults {
  if (!item) {
    return result;
  }

  if (hasItem(result, item)) {
    return result;
  }

  const cloned = cloneDeep(result);
  const total = result.total + 1;
  cloned.results.unshift(item);
  const results = cloned.results;
  const [first, last] = getBoundaries(results);

  return {
    ...cloned,
    first,
    last,
    total,
    results,
  };
}

export function removeItem(result: FeedResults, item: FeedItem): FeedResults {
  if (!item) {
    return result;
  }

  if (!hasItem(result, item)) {
    return result;
  }

  const cloned = cloneDeep(result);
  const total = result.total - 1;
  const results = cloned.results.filter((i) => i.id !== item.id);
  const [first, last] = getBoundaries(results);

  return {
    ...cloned,
    first,
    last,
    total,
    results,
  };
}
export function removeItemWithId(result: FeedResults, id: string): FeedResults {
  if (!id) {
    return result;
  }

  const index = result.results.findIndex((item) => item.id === id);
  if (index < 0) {
    return result;
  }

  const cloned = cloneDeep(result);
  const total = result.total - 1;
  const results = cloned.results.filter((item) => item.id !== id);
  const [first, last] = getBoundaries(results);

  return {
    ...cloned,
    first,
    last,
    total,
    results,
  };
}
