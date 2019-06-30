export interface Source {
  id: string;
  url: string;
  title: string;
  type: string;
  status: string;
  isPaused: boolean;
  createdAt: string;
  updatedAt: string;
  parsedAt: string;
  stats?: Stats;
}

export interface SyndicationResults {
  limit: number;
  total: number;
  offset: number;
  results: Source[];
}

export interface SearchParams {
  isPaused: boolean;
}

export interface Stats {
  statusCode: number;
  frequency: string;
  totalEntries: number;
  totalSuccess: number;
  totalFailure: number;
}
