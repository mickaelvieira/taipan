import { cloneDeep } from "lodash";
import { SyndicationResults, Source } from "../../../types/syndication";

export function removeSource(
  result: SyndicationResults,
  item: Source
): SyndicationResults {
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

  return {
    ...cloned,
    total,
    results
  };
}
