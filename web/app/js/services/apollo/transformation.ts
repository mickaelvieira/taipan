/* eslint @typescript-eslint/no-explicit-any: "off" */
/* eslint @typescript-eslint/no-use-before-define: "off" */

export interface WithTypename {
  __typename: string;
}

type Item = Record<string, any>;

interface Collection {
  results: Item[];
}

interface Transformers {
  [type: string]: (data: Item | Collection) => Item | Collection;
}

const transformers: Transformers = {
  Document: transformItem,
  Bookmark: transformItem,
  Source: transformItem,
  UserSubscription: transformItem,
  Image: transformItem,
  User: transformItem,
  Email: transformItem,
  FeedBookmarkResults: (data) => transformCollection(data as Collection),
  FeedDocumentResults: (data) => transformCollection(data as Collection),
  BookmarkSearchResults: (data) => transformCollection(data as Collection),
  DocumentSearchResults: (data) => transformCollection(data as Collection),
  SubscriptionCollection: (data) => transformCollection(data as Collection),
  SourceCollection: (data) => transformCollection(data as Collection),
};

const isObject = (value: any): boolean =>
  value && typeof value === "object" && value.constructor === Object;
const isURL = (name: string): boolean => ["url", "domain"].includes(name);
const isDate = (name: string): boolean =>
  ["createdAt", "updatedAt", "addedAt", "parsedAt", "favoritedAt"].includes(
    name
  );

const isTransformable = (type?: string): boolean =>
  !!type && Object.keys(transformers).includes(type);

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
  if (input && isTransformable(input.__typename)) {
    const transformer = transformers[input.__typename];
    return transformer(input);
  }
  return input;
}
