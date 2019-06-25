import { cloneDeep } from "lodash";
import {
  FeedQueryData,
  FeedEventData,
  FeedActions,
  FeedAction,
  FeedResults,
  FeedItem,
  FeedEvent
} from "../../../types/feed";

export function getDataKey(data: FeedQueryData): string | null {
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
    limit: 0
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

function addItemToFeedResults(
  result: FeedResults,
  item: FeedItem
): FeedResults {
  if (!item) {
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
    results
  };
}

function removeItemFromFeedResults(
  result: FeedResults,
  item: FeedItem
): FeedResults {
  if (!item) {
    return result;
  }

  const index = result.results.findIndex(i => i.id === item.id);
  if (index < 0) {
    return result;
  }

  const cloned = cloneDeep(result);
  const total = result.total - 1;
  const results = cloned.results.filter(i => i.id !== item.id);
  const [first, last] = getBoundaries(results);

  return {
    ...cloned,
    first,
    last,
    total,
    results
  };
}

export const feedResultsAction: FeedActions<FeedAction> = {
  Add: addItemToFeedResults,
  Remove: removeItemFromFeedResults
};
