import { Datetime } from "./scalars";

export interface HTTPClientLog {
  id: string;
  checksum: string;
  requestURI: string;
  statusCode: number;
  contentType: string;
  requestMethod: string;
  hasFailed: boolean;
  failureReason: string;
  createdAt: Datetime;
}
