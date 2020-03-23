import { User } from "../../types/users";
import { StoreName, Mode } from "../db/types";
import { getDBStore } from "../idb";

/* eslint @typescript-eslint/no-explicit-any: off */
export async function upsert(item: User): Promise<any> {
  const store = await getDBStore(StoreName.USER, Mode.READWRITE);
  return store.upsert(item);
}

export async function get(): Promise<User | null> {
  const store = await getDBStore(StoreName.USER);
  const results = await store.select();
  return results.length > 0 ? results[0] : null;
}

export async function remove(id: string): Promise<any> {
  const store = await getDBStore(StoreName.USER, Mode.READWRITE);
  return store.delete(id);
}

export default {
  upsert,
  delete: remove,
  get,
};
