import { Source } from "../types/syndication";
import { Subscription } from "../types/subscription";

export function getDomain(item: Source | Subscription): URL {
  return item.domain ? item.domain : item.url;
}
