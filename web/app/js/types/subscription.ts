import { Datetime } from "./scalars";

export interface Subscription {
  id: string;
  url: string;
  title: string;
  type: string;
  isSubscribed: boolean;
  createdAt: Datetime;
  updatedAt: Datetime;
}

export interface SubscriptionResults {
  limit: number;
  total: number;
  offset: number;
  results: Subscription[];
}
