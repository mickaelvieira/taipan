import { Query } from "react-apollo";
import query from "../graphql/query/bot/logs.graphql";
import { OffsetPagination } from "../../../types";
import { LogResults } from "../../../types/http";

export interface Data {
  bot: {
    logs: LogResults;
  };
}

export interface Variables {
  url: string;
  pagination: OffsetPagination;
}

const variables = {
  pagination: {
    limit: 50
  }
};

export { query, variables };

class LogsQuery extends Query<Data, Variables> {
  static defaultProps = {
    query,
    variables
  };
}

export default LogsQuery;
