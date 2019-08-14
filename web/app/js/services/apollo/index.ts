import ApolloClient from "apollo-client";
import { InMemoryCache, IdGetterObj } from "apollo-cache-inmemory";
// import { ApolloLink, split, concat } from "apollo-link";
import { ApolloLink, concat } from "apollo-link";
import { onError } from "apollo-link-error";
import { HttpLink } from "apollo-link-http";
import { setContext } from "apollo-link-context";
// import { WebSocketLink } from "apollo-link-ws";
// import { getMainDefinition } from "apollo-utilities";
import { Log } from "../../types/http";
import { AppInfo } from "../../types/app";
import { Image } from "../../types/image";
import { User } from "../../types/users";
import { Email } from "../../types/users";
import { Document } from "../../types/document";
import { Bookmark } from "../../types/bookmark";
import { Subscription } from "../../types/subscription";
import { Source } from "../../types/syndication";
import transform from "./transform";

function isSubscriptionOperation(name: string): boolean {
  return ["onBookmarkChange", "onDocumentChange", "onUserChange"].includes(
    name
  );
}

export function genRandomId(): string {
  return [...Array(20)]
    .map(() => (~~(Math.random() * 36)).toString(36))
    .join("");
}

type Models = IdGetterObj &
  Partial<AppInfo> &
  Partial<User> &
  Partial<Email> &
  Partial<Document> &
  Partial<Bookmark> &
  Partial<Image> &
  Partial<Source> &
  Partial<Subscription> &
  Partial<Log>;

export default (clientId: string): ApolloClient<object> => {
    const cache = new InMemoryCache({
    freezeResults: true,
    dataIdFromObject: ({ id, url, name, value, __typename }: Models) => {
      if (!__typename) {
        return null;
      }
      if (id) {
        return `${__typename}@${id}`;
      }
      if (url) {
        return `${__typename}@${url}`;
      }
      if (__typename === "Email") {
        return `${__typename}@${value}`;
      }
      if (__typename === "AppInfo") {
        return `${__typename}@${name}`;
      }
      return null;
    }
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
  // const wsLink = new WebSocketLink({
  //   uri: `ws${isEncrypted ? "s" : ""}:${endpoint}`,
  //   options: {
  //     reconnect: true,
  //     reconnectionAttempts: Infinity,
  //     inactivityTimeout: 0
  //   }
  // });
  // const transportLink = split(
  //   ({ query }) => {
  //     const definition = getMainDefinition(query);
  //     return (
  //       definition.kind === "OperationDefinition" &&
  //       definition.operation === "subscription"
  //     );
  //   },
  //   wsLink,
  //   httpLink
  // );
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
  const transformationLink = new ApolloLink((operation, forward) => {
    const name = operation.operationName;
    if (!forward) {
      return null;
    }
    if (isSubscriptionOperation(name)) {
      return forward(operation);
    }
    return forward(operation).map(transform);
  });
  const clientIdLink = setContext(() => ({
    headers: {
      [headerName]: clientId
    }
  }));

  const link = concat(
    concat(concat(errorLink, clientIdLink), transformationLink),
    httpLink
  );
  const client = new ApolloClient({
    link,
    cache
  });

  return client;
};
