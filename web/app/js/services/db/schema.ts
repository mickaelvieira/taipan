import { KeyPaths, StoreName } from "./types";

export default function(db: IDBDatabase): void {
  const bookmarks = db.createObjectStore(StoreName.BOOKMARKS, {
    keyPath: KeyPaths.ID
  });
  bookmarks.createIndex("hash", "hash", { unique: true });
  bookmarks.createIndex("href", "href", { unique: true });

  db.createObjectStore(StoreName.USER, {
    keyPath: KeyPaths.ID
  });

  db.createObjectStore(StoreName.FEED, {
    keyPath: KeyPaths.TYPE
  });
}
