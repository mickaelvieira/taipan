import { KeyPaths, StoreName, Mode } from "../db/types";
import { getDBStore } from "../idb";

export enum FeedTypes {
  LATEST = "latest"
}

interface Feed {
  type: FeedTypes;
  results: string[];
}

export async function upsert(item: Feed) {
  const store = await getDBStore(StoreName.FEED, Mode.READWRITE);
  return store.upsert(item, KeyPaths.TYPE);
}

export async function update(item: Feed) {
  const store = await getDBStore(StoreName.FEED, Mode.READWRITE);
  return store.update(item);
}

export async function batch(items: Feed[]) {
  const store = await getDBStore(StoreName.FEED, Mode.READWRITE);
  return store.batch(items);
}

export async function all(): Promise<Feed> {
  const store = await getDBStore(StoreName.FEED);
  const results = await store.select();
  return {
    type: FeedTypes.LATEST,
    results
  };
}

export default {
  upsert,
  update,
  batch,
  all
};
