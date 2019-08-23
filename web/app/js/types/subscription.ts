import { Tag } from "./syndication";

export interface Subscription {
  id: string;
  url: URL;
  domain: URL | null;
  title: string;
  type: string;
  isSubscribed: boolean;
  frequency: string;
  createdAt: Date | null;
  updatedAt: Date | null;
}

export interface SubscriptionResults {
  limit: number;
  total: number;
  offset: number;
  results: Subscription[];
}

export interface SearchParams {
  terms: string[];
  tags: string[];
}

export interface TagResults {
  total: number;
  results: Tag[];
}
