import ApolloClient from "apollo-client";
import { IdGetterObj } from "apollo-boost";
import { InMemoryCache } from "apollo-cache-inmemory";
import { split } from "apollo-link";
import { HttpLink } from "apollo-link-http";
import { WebSocketLink } from "apollo-link-ws";
import { getMainDefinition } from "apollo-utilities";

export default () => {
  const cache = new InMemoryCache({
    freezeResults: true,
    dataIdFromObject: ({ id, __typename }: IdGetterObj) =>
      id && __typename ? `${__typename}@${id}` : null
  });

  const endpoint =
    process.env.NODE_ENV === "production"
      ? process.env.APP_GRAPHQL_ENDPOINT
      : "//localhost:9000/graphql";

  const protocol = process.env.NODE_ENV === "production" ? "https:" : "http:";

  const httpLink = new HttpLink({
    uri: `${protocol}${endpoint}`
  });
  const wsLink = new WebSocketLink({
    uri: `ws:${endpoint}`,
    options: {
      reconnect: true
    }
  });

  const link = split(
    ({ query }) => {
      const definition = getMainDefinition(query);
      return (
        definition.kind === "OperationDefinition" &&
        definition.operation === "subscription"
      );
    },
    wsLink,
    httpLink
  );

  const client = new ApolloClient({
    link,
    cache
  });

  return client;
};
