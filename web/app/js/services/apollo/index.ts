import ApolloClient from "apollo-boost";

export default () => {
  const client = new ApolloClient({
    uri: "/graphql"
  });

  return client;
};
