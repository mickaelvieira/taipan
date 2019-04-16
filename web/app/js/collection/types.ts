import { RootState } from "store/reducer/default";

export interface Error {
  title: string;
  code: string;
  message: string;
}

export interface Data {
  name: string;
  value: string | number | boolean | null;
}

export interface Link {
  href: string;
  rel: string;
  render: string;
}

export interface Item {
  href: string;
  data: Data[];
  links: Link[];
}

export interface Template extends Object {
  data: Data[];
}

export interface Query {
  href: string;
  rel: string;
  prompt: string;
  data: Data[];
}

export interface Collection {
  version: string;
  href: string;
  items: Item[];
  links: Link[];
  template?: Template;
  queries?: Query[];
  error?: Error;
}

export interface CollectionResponse extends Object {
  collection: Collection;
}

export interface DBStateResponse extends RootState {}
