import { cloneDeep } from "lodash";
import {
  FeedActions,
  FeedAction,
  FeedResults,
  FeedItem
} from "../../../types/feed";

function getBoundaries(results: FeedItem[]): [string, string] {
  let first = "";
  let last = "";
  if (results.length > 0) {
    first = results[0].id;
    last = results[results.length - 1].id;
  }
  return [first, last];
}

export function addItemsFromFeedResults(
  result: FeedResults,
  items: FeedResults
): [FeedResults, FeedResults] {
  const cloned = cloneDeep(result);
  const total = result.total + items.results.length;
  cloned.results.unshift(...items.results);

  const results = cloned.results;
  const [first, last] = getBoundaries(results);

  const newsResults = {
    ...cloned,
    first,
    last,
    total,
    results
  };

  const latestNewsResults = {
    ...items,
    results: []
  };

  return [newsResults, latestNewsResults];
}

function addItemToFeedResults(
  result: FeedResults,
  item: FeedItem
): FeedResults {
  // @TODO I might have to add a condition to check whether the item is already in the cache
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
