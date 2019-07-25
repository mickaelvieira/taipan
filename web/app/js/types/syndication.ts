import { Datetime } from "./scalars";
import { HTTPClientLog } from "./http";

export interface Source {
  id: string;
  url: string;
  domain: string | null;
  title: string;
  type: string;
  status: string;
  isPaused: boolean;
  isDeleted: boolean;
  frequency: string;
  createdAt: Datetime;
  updatedAt: Datetime;
  parsedAt: Datetime | null;
  stats?: Stats;
  logEntries?: HTTPClientLog[];
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
