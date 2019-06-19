import ApolloClient from "apollo-client";
import { IdGetterObj } from "apollo-boost";
import { InMemoryCache } from "apollo-cache-inmemory";
import { split, concat } from "apollo-link";
import { onError } from "apollo-link-error";
import { HttpLink } from "apollo-link-http";
import { WebSocketLink } from "apollo-link-ws";
import { getMainDefinition } from "apollo-utilities";

export default () => {
  const cache = new InMemoryCache({
    freezeResults: true,
    dataIdFromObject: ({ id, __typename }: IdGetterObj) =>
      id && __typename ? `${__typename}@${id}` : null
  });

  const isEncrypted = process.env.APP_GRAPHQL_ENCRYPTED === "true";
  const endpoint = process.env.APP_GRAPHQL_ENDPOINT;

  const httpLink = new HttpLink({
    uri: `http${isEncrypted ? "s" : ""}:${endpoint}`
  });
  const wsLink = new WebSocketLink({
    uri: `ws${isEncrypted ? "s" : ""}:${endpoint}`,
    options: {
      reconnect: true
    }
  });

  const transportLink = split(
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

  const errorLink = onError(({ graphQLErrors, networkError }) => {
    if (graphQLErrors)
      graphQLErrors.map(({ message, locations, path }) =>
        console.log(
          `[GraphQL error]: Message: ${message}, Location: ${locations}, Path: ${path}`
        )
      );
    if (networkError) {
      console.log(`[Network error]: ${networkError}`);
    }
  });

  const link = concat(errorLink, transportLink);
  const client = new ApolloClient({
    link,
    cache
  });

  return client;
};
