import Store from "./db/store";
import { Mode, StoreName, DBConfig, DBUpdater } from "./db/types";

/* eslint @typescript-eslint/no-explicit-any: off */

let config = new Map();

function getDB(
  config: Map<string, any>,
  onUpgradeNeeded?: DBUpdater
): Promise<IDBDatabase> {
  return new Promise<IDBDatabase>((resolve, reject) => {
    if (!config.get("name") || !config.get("version")) {
      reject(new Error("Configuration is missing"));
      return;
    }

    const request = window.indexedDB.open(
      config.get("name"),
      config.get("version")
    );

    request.onerror = () => reject(request.error);
    request.onsuccess = () => resolve(request.result);
    request.onupgradeneeded = () => {
      const db = request.result;
      db.onerror = () => reject(db.error);
      if (typeof onUpgradeNeeded === "function") {
        onUpgradeNeeded(db);
      }
    };
  });
}

export async function getDBStore(
  name: StoreName,
  mode: Mode = Mode.READONLY
): Promise<Store> {
  const db = await getDB(config);
  const tr = db.transaction([name], mode);

  return new Store(tr.objectStore(name));
}

export default async function initDBStore(
  cnf: DBConfig,
  onUpgrade: DBUpdater
): Promise<IDBDatabase> {
  if (Object.keys(cnf).length === 0) {
    throw new Error("You need to provide a configuration");
  }

  config = new Map(Object.entries(cnf));

  return await getDB(config, onUpgrade);
}
