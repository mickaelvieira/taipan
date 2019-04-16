import { Bookmark } from "types/bookmark";
import { ResultsFilter, StoreName, Mode } from "services/db/types";
import { getDBStore } from "../idb";

export async function upsert(item: Bookmark) {
  const store = await getDBStore(StoreName.BOOKMARKS, Mode.READWRITE);
  return store.upsert(item);
}

export async function batch(items: Bookmark[]) {
  const store = await getDBStore(StoreName.BOOKMARKS, Mode.READWRITE);
  return store.batch(items);
}

export async function all(): Promise<Bookmark[]> {
  const store = await getDBStore(StoreName.BOOKMARKS);
  return store.select();
}

export async function unread(): Promise<Bookmark[]> {
  const store = await getDBStore(StoreName.BOOKMARKS);
  return store.select((item: Bookmark) => !item.is_read);
}

export async function select(filter: ResultsFilter): Promise<Bookmark[]> {
  const store = await getDBStore(StoreName.BOOKMARKS);
  return store.select(filter);
}

export async function remove(id: string): Promise<{}> {
  const store = await getDBStore(StoreName.BOOKMARKS, Mode.READWRITE);
  return store.delete(id);
}

export default {
  upsert,
  batch,
  all,
  select,
  unread,
  delete: remove
};
