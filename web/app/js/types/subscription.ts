import { Datetime } from "./scalars";

export interface Subscription {
  id: string;
  url: string;
  domain: string | null;
  title: string;
  type: string;
  isSubscribed: boolean;
  frequency: string;
  createdAt: Datetime | null;
  updatedAt: Datetime | null;
}

export interface SubscriptionResults {
  limit: number;
  total: number;
  offset: number;
  results: Subscription[];
}

export interface SearchParams {
  terms: string[];
}
