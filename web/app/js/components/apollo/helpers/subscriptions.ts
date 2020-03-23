import { cloneDeep } from "lodash";
import { SubscriptionResults, Subscription } from "../../../types/subscription";

function hasItem(result: SubscriptionResults, item: Subscription): boolean {
  const index = result.results.findIndex(({ id }) => id === item.id);
  return index >= 0;
}

export function addSubscription(
  result: SubscriptionResults,
  item: Subscription
): SubscriptionResults {
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
    results,
  };
}

export function removeSubscription(
  result: SubscriptionResults,
  item: Subscription
): SubscriptionResults {
  if (!item) {
    return result;
  }

  if (!hasItem(result, item)) {
    return result;
  }

  const cloned = cloneDeep(result);
  const total = result.total - 1;
  const results = cloned.results.filter((i) => i.id !== item.id);

  return {
    ...cloned,
    total,
    results,
  };
}
