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
