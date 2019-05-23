import { Query } from "react-apollo";
import { Bookmark } from "../../../types/bookmark";
import query from "../../../services/apollo/query/favorites.graphql";

export interface Data {
  GetFavorites: {
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

class FavoritesQuery extends Query<Data, Variables> {}

export default FavoritesQuery;
