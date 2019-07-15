import ApolloClient from "apollo-client";
import { IdGetterObj } from "apollo-boost";
import { InMemoryCache } from "apollo-cache-inmemory";
import { split, concat } from "apollo-link";
import { onError } from "apollo-link-error";
import { HttpLink } from "apollo-link-http";
import { setContext } from "apollo-link-context";
import { WebSocketLink } from "apollo-link-ws";
import { getMainDefinition } from "apollo-utilities";

export function genRandomId(): string {
  return [...Array(20)]
    .map(() => (~~(Math.random() * 36)).toString(36))
    .join("");
}

export default (clientId: string) => {
  const cache = new InMemoryCache({
    freezeResults: true,
    dataIdFromObject: ({ id, __typename }: IdGetterObj) =>
      id && __typename ? `${__typename}@${id}` : null
  });

  const headerName = process.env.APP_CLIENT_ID_HEADER as string;
  const isEncrypted = !!process.env.APP_GRAPHQL_ENCRYPTED;
  const endpoint = process.env.APP_GRAPHQL_ENDPOINT;

  if (!headerName) {
    throw new Error("Client ID header name must be defined.");
  }

  const httpLink = new HttpLink({
    uri: `http${isEncrypted ? "s" : ""}:${endpoint}`
  });
  const wsLink = new WebSocketLink({
    uri: `ws${isEncrypted ? "s" : ""}:${endpoint}`,
    options: {
      reconnect: true,
      reconnectionAttempts: Infinity,
      connectionCallback: _ => {
        // console.log("Connected to WS");
        // console.log(_);
      },
      inactivityTimeout: 0
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

  const clientIdLink = setContext(() => ({
    headers: {
      [headerName]: clientId
    }
  }));

  const link = concat(concat(errorLink, clientIdLink), transportLink);
  const client = new ApolloClient({
    link,
    cache
  });

  return client;
};
