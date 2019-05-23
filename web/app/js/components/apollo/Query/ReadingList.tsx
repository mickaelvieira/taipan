import { Query } from "react-apollo";
import { Bookmark } from "../../../types/bookmark";
import query from "../../../services/apollo/query/reading-list.graphql";

export interface Data {
  GetReadingList: {
    total: number;
    offset: number;
    limit: number;
    results: Bookmark[];
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

class ReadingListQuery extends Query<Data, Variables> {}

export default ReadingListQuery;
