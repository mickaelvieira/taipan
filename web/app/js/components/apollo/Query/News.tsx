import { Query } from "react-apollo";
import { Document } from "../../../types/document";
import query from "../../../services/apollo/query/news.graphql";

export interface Data {
  News: {
    total: number;
    offset: number;
    limit: number;
    results: Document[];
  };
}

export interface Variables {
  offset?: number;
  limit?: number;
}

const variables = {
  limit: 10
};

export { query, variables };

class NewsQuery extends Query<Data, Variables> {}

export default NewsQuery;
