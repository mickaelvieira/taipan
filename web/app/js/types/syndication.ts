import { Log } from "./http";

export interface Source {
  id: string;
  url: URL;
  domain: URL | null;
  title: string;
  type: string;
  isPaused: boolean;
  isDeleted: boolean;
  frequency: string;
  createdAt: Date;
  updatedAt: Date;
  parsedAt: Date | null;
  stats?: Stats;
  logEntries?: Log[];
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
  subscribers: number;
  totalEntries: number;
  totalSuccess: number;
  totalFailure: number;
}
