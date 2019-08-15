import { FetchResult } from "apollo-link";
import transformation, { WithTypename } from "./transformation";

/* eslint @typescript-eslint/no-explicit-any: "off" */

export default function(fetchResults: FetchResult): FetchResult {
  const data: Record<string, any> = {};

  for (const [ns, op] of Object.entries(fetchResults.data)) {
    if (op) {
      const { __typename, ...rest } = op as Record<string, any> & WithTypename;
      const keys = Object.keys(rest);
      const key = keys.length ? keys[0] : null;

      if (key) {
        const result = rest[key] as Record<string, any>;
        data[ns] = {
          __typename,
          [key]: transformation(result)
        };
      }
    } else {
      data[ns] = op;
    }
  }

  return {
    ...fetchResults,
    data
  };
}
