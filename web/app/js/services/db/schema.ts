import { KeyPaths, StoreName } from "./types";

export default function(db: IDBDatabase) {
  const bookmarks = db.createObjectStore(StoreName.BOOKMARKS, {
    keyPath: KeyPaths.ID
  });
  bookmarks.createIndex("hash", "hash", { unique: true });
  bookmarks.createIndex("href", "href", { unique: true });

  const user = db.createObjectStore(StoreName.USER, {
    keyPath: KeyPaths.ID
  });

  const feed = db.createObjectStore(StoreName.FEED, {
    keyPath: KeyPaths.TYPE
  });
}
