import { toPromisedCursor, toPromisedRequest } from "./helpers";
import { BatchOperation, KeyPaths, ResultsFilter } from "./types";

type Process = (callback: BatchOperation) => Promise<any>;

/* eslint @typescript-eslint/no-explicit-any: off */
function batch(items: any[]): Process {
  const process = (callback: BatchOperation): Promise<any> => {
    const item = items.shift();
    return item
      ? callback(item).then(() => process(callback))
      : Promise.resolve();
  };
  return process;
}

/* eslint @typescript-eslint/explicit-function-return-type: off */
/* eslint @typescript-eslint/explicit-module-boundary-types: off */
export default class Store {
  store: IDBObjectStore;

  constructor(store: IDBObjectStore) {
    this.store = store;
  }

  create = (item: any) => {
    return toPromisedRequest(this.store.add(item));
  };

  read = (id: string) => {
    return toPromisedRequest(this.store.get(id));
  };

  update = (item: any) => {
    return toPromisedRequest(this.store.put(item));
  };

  upsert = async (item: any, keyPath: string = KeyPaths.ID) => {
    const existing = await this.read(item[keyPath]);
    return existing ? this.update(item) : this.create(item);
  };

  delete = (id: string) => {
    return toPromisedRequest(this.store.delete(id));
  };

  select = (filter: ResultsFilter = (item: any) => item) => {
    return toPromisedCursor(this.store.openCursor(), filter);
  };

  batch = (items: any[], op = this.upsert) => {
    return batch(items.slice())(op);
  };
}
