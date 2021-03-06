import { ResultsFilter } from "./types";

/* eslint @typescript-eslint/no-explicit-any: off */
export function toPromisedRequest(request: IDBRequest): Promise<any> {
  return new Promise((resolve, reject) => {
    request.onsuccess = () => resolve(request.result);
    request.onerror = () => reject(request.error);
  });
}

export function toPromisedCursor(
  request: IDBRequest,
  filter: ResultsFilter
): Promise<any[]> {
  return new Promise((resolve, reject) => {
    const results: any[] = [];
    request.onsuccess = (event) => {
      const cursor = event.target && event.target.result;
      if (cursor) {
        const item = cursor.value;
        if (filter(item)) {
          results.push(item);
        }
        cursor.continue();
      } else {
        resolve(results);
      }
    };
    request.onerror = () => reject(request.error);
  });
}
