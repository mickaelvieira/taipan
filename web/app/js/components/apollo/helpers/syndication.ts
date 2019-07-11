import { cloneDeep } from "lodash";
import { SyndicationResults, Source } from "../../../types/syndication";

function hasItem(result: SyndicationResults, item: Source): boolean {
  const index = result.results.findIndex(({ id }) => id === item.id);
  return index >= 0;
}

export function addSource(
  result: SyndicationResults,
  item: Source
): SyndicationResults {
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

  return {
    ...cloned,
    total,
    results
  };
}

export function removeSource(
  result: SyndicationResults,
  item: Source
): SyndicationResults {
  if (!item) {
    return result;
  }

  if (!hasItem(result, item)) {
    return result;
  }

  const cloned = cloneDeep(result);
  const total = result.total - 1;
  const results = cloned.results.filter(i => i.id !== item.id);

  return {
    ...cloned,
    total,
    results
  };
}
