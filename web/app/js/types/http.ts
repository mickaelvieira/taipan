import { Datetime } from "./scalars";

export interface HTTPClientLog {
  id: string;
  checksum: string;
  requestURI: string;
  statusCode: number;
  contentType: string;
  createdAt: Datetime;
}
