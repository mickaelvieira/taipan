import { Source } from "../types/syndication";
import { Subscription } from "../types/subscription";

export function getDomain(item: Source | Subscription): URL {
  return item.domain ? new URL(item.domain) : new URL(item.url);
}
