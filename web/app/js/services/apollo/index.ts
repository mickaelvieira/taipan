import ApolloClient, { IdGetterObj } from "apollo-boost";
import { InMemoryCache } from "apollo-cache-inmemory";

export default () => {
  const cache = new InMemoryCache({
    freezeResults: true,
    dataIdFromObject: ({ id, __typename }: IdGetterObj) =>
      id && __typename ? `${__typename}@${id}` : null
  });

  const client = new ApolloClient({
    uri: "/graphql",
    cache
  });

  return client;
};
