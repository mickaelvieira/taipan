export interface Subscription {
  id: string;
  url: string;
  title: string;
  type: string;
  isSubscribed: boolean;
  createdAt: string;
  updatedAt: string;
}

export interface SubscriptionResults {
  limit: number;
  total: number;
  offset: number;
  results: Subscription[];
}
