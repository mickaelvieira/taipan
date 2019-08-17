import query from "../graphql/query/bot/logs.graphql";
import { OffsetPagination } from "../../../types";
import { LogResults } from "../../../types/http";

export interface Data {
  bot: {
    logs: LogResults;
  };
}

export interface Variables {
  url: URL;
  pagination: OffsetPagination;
}

const variables = {
  pagination: {
    limit: 50
  }
};

export { query, variables };
