import { Bookmark } from "../../types/bookmark";
import { ResultsFilter, StoreName, Mode } from "../db/types";
import { getDBStore } from "../idb";

/* eslint @typescript-eslint/no-explicit-any: off */
export async function upsert(item: Bookmark): Promise<any> {
  const store = await getDBStore(StoreName.BOOKMARKS, Mode.READWRITE);
  return store.upsert(item);
}

export async function batch(items: Bookmark[]): Promise<any> {
  const store = await getDBStore(StoreName.BOOKMARKS, Mode.READWRITE);
  return store.batch(items);
}

export async function all(): Promise<Bookmark[]> {
  const store = await getDBStore(StoreName.BOOKMARKS);
  return store.select();
}

export async function unfavorite(): Promise<Bookmark[]> {
  const store = await getDBStore(StoreName.BOOKMARKS);
  return store.select((item: Bookmark) => !item.isFavorite);
}

export async function select(filter: ResultsFilter): Promise<Bookmark[]> {
  const store = await getDBStore(StoreName.BOOKMARKS);
  return store.select(filter);
}

/* eslint @typescript-eslint/no-explicit-any: off */
export async function remove(id: string): Promise<any> {
  const store = await getDBStore(StoreName.BOOKMARKS, Mode.READWRITE);
  return store.delete(id);
}

export default {
  upsert,
  batch,
  all,
  select,
  unfavorite,
  delete: remove,
};
