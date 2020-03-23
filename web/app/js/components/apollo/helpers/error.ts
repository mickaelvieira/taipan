import { ApolloError } from "apollo-client";

export function getErrorMessage({
  graphQLErrors,
  networkError,
}: ApolloError): string {
  if (graphQLErrors && graphQLErrors.length > 0) {
    return graphQLErrors[0].message;
  }
  if (networkError) {
    return networkError.message;
  }
  return "";
}
