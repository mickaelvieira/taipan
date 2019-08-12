export interface Log {
  id: string;
  checksum: string;
  requestURI: string;
  statusCode: number;
  contentType: string;
  requestMethod: string;
  hasFailed: boolean;
  failureReason: string;
  createdAt: Date;
}

export interface LogResults {
  limit: number;
  total: number;
  offset: number;
  results: Log[];
}
