/* eslint @typescript-eslint/no-explicit-any: "off" */
/* eslint @typescript-eslint/no-use-before-define: "off" */

export interface WithTypename {
  __typename: string;
}

type Item = Record<string, any>;

interface Collection {
  results: Item[];
}

const types = {
  url: ["url", "domain"],
  date: ["createdAt", "updatedAt", "addedAt", "parsedAt", "favoritedAt"],
};

const isObject = (value: any): boolean =>
  value && typeof value === "object" && value.constructor === Object;
const isURL = (name: string): boolean => types.url.includes(name);
const isDate = (name: string): boolean => types.date.includes(name);
const isCollection = (entity: any): boolean =>
  typeof entity.results !== "undefined";

function transformItem(input: Item): Item {
  const output: Record<string, any> = {};
  for (const [key, value] of Object.entries(input)) {
    if (value) {
      if (isURL(key)) {
        output[key] = new URL(value);
      } else if (isDate(key)) {
        output[key] = new Date(value);
      } else if (isObject(value)) {
        output[key] = transform(value);
      } else if (Array.isArray(value)) {
        output[key] = value.map(transform);
      } else {
        output[key] = value;
      }
    } else {
      output[key] = value;
    }
  }
  return output as Item;
}

function transformCollection(result: Collection): Collection {
  return {
    ...result,
    results: result.results.map(transform),
  };
}

export default function transform(
  input: Record<string, any>
): Item | Collection {
  if (input) {
    if (isCollection(input)) {
      return transformCollection(input as Collection);
    }
    return transformItem(input);
  }
  return input;
}
