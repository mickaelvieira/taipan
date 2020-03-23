/* eslint @typescript-eslint/no-explicit-any: off */
export enum KeyPaths {
  ID = "id",
  TYPE = "type",
}

export enum StoreName {
  BOOKMARKS = "bookmarks",
  FEED = "feed",
  USER = "user",
}

export enum Mode {
  READONLY = "readonly",
  READWRITE = "readwrite",
  VERSIONCHANGE = "versionchange",
}

export interface BatchOperation {
  (item: any): Promise<any>;
}

export interface ResultsFilter {
  (item: any): boolean;
}

export interface DBConfig {
  name: string;
  version: number;
}

export interface DBUpdater {
  (db: IDBDatabase): void;
}
