import { BookmarkHistory, BookmarkHistoryEntry } from "types/bookmark";

type Entries = Array<BookmarkHistoryEntry>;

const sortEntriesByCreationDate = (items: Entries): Entries =>
  items.sort(
    (a, b): -1 | 0 | 1 => {
      const d1 = new Date(a.data.created_at);
      const d2 = new Date(b.data.created_at);

      if (d1 < d2) {
        return -1;
      }

      if (d1 > d2) {
        return 1;
      }

      return 0;
    }
  );

const getFailedEntries = (items: Entries) =>
  items.filter(item => item.data.response_code >= 400);

const getSuccessfulEntries = (items: Entries) =>
  items.filter(item => item.data.response_code < 400);

export function getLastFailure(
  history: BookmarkHistory
): BookmarkHistoryEntry | null | undefined {
  const items = sortEntriesByCreationDate(getFailedEntries(history.items));
  return items.length > 0 ? items.pop() : null;
}

export function getLastSuccess(
  history: BookmarkHistory
): BookmarkHistoryEntry | null | undefined {
  const items = sortEntriesByCreationDate(getSuccessfulEntries(history.items));
  return items.length > 0 ? items.pop() : null;
}

export function getLastEntry(
  history: BookmarkHistory
): BookmarkHistoryEntry | null | undefined {
  const items = sortEntriesByCreationDate(history.items);
  return items.length > 0 ? items.pop() : null;
}
