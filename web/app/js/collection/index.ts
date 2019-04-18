import { Collection, Data, Link, Item } from "./types";
import { Bookmark, Bookmarks } from "../types/bookmark";

export const flattenLinks = (links: Link[]) =>
  links.reduce((carry, { rel, href }) => ({ ...carry, [rel]: href }), {});

export const flattenData = (data: Data[]) =>
  data.reduce((carry, { name, value }) => ({ ...carry, [name]: value }), {});

export const parseItem = (item: Item) => {
  const href = { item };
  const data = "data" in item ? flattenData(item.data) : {};
  const links = "links" in item ? flattenLinks(item.links) : {};

  return {
    ...data,
    href,
    links
  };
};

export const parseItems = (items: Item[]) => items.map(item => parseItem(item));

export const parseCollection = (collection: Collection) => {
  const { href } = collection;
  const links = "links" in collection ? flattenLinks(collection.links) : {};
  const items = "items" in collection ? parseItems(collection.items) : [];

  return {
    href,
    links,
    items
  };
};

export const getAllIds = (collection: Bookmarks) =>
  collection.items.map(({ id }) => id);

export const arrayToObject = (items: Bookmark[]) =>
  items.reduce((carry, item) => ({ ...carry, [item.id]: item }), {});

export const arrayToMap = (items: Bookmark[]) =>
  new Map<string, any>(items.map(item => [item.id, item]));
