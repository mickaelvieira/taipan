import { Tag } from "../types/syndication";

export function sort(tags: Tag[]): Tag[] {
  const sorted = [...tags];

  sorted.sort(function(a, b) {
    if (a.label < b.label) {
      return -1;
    }
    if (a.label > b.label) {
      return 1;
    }
    return 0;
  });

  return sorted;
}
