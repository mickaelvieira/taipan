import { Query } from "react-apollo";
import query from "../graphql/query/subscriptions/subscriptions.graphql";
import { OffsetPagination } from "../../../types";
import { SubscriptionResults } from "../../../types/subscription";

export interface Data {
  subscriptions: {
    subscriptions: SubscriptionResults;
  };
}

export interface Variables {
  pagination: OffsetPagination;
}

const variables = {
  pagination: {
    limit: 50
  }
};

export { query, variables };

class SubscriptionsQuery extends Query<Data, Variables> {
  static defaultProps = {
    query,
    variables
  };
}

export default SubscriptionsQuery;
