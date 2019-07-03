import { Pact, ApolloGraphQLInteraction } from "@pact-foundation/pact"
import ApolloClient from "apollo-client";
import { IdGetterObj } from "apollo-boost";
import { HttpLink } from "apollo-link-http";
import { InMemoryCache } from "apollo-cache-inmemory";
import { query } from "./Init"
// import query from "../../../services/apollo/query/users/";

describe("query hello on /graphql", () => {
  let client: ApolloClient<object>;
  let provider: Pact;
  beforeAll(() => {
    const link = new HttpLink({
      uri: `http://localhost:9001`
    });
    const cache = new InMemoryCache({
      freezeResults: true,
      dataIdFromObject: ({ id, __typename }: IdGetterObj) =>
        id && __typename ? `${__typename}@${id}` : null
    });
    client = new ApolloClient({
      link,
      cache
    });

    provider = new Pact({
      port: 9001,
      // log: path.resolve(process.cwd(), "logs", "mockserver-integration.log"),
      // dir: path.resolve(process.cwd(), "pacts"),
      // spec: 2,
      // cors: true,
      // pactfileWriteMode: "update",
      consumer: "GraphQLConsumer",
      provider: "GraphQLProvider"
    });
  });

  beforeEach(() => {
    const graphqlQuery = new ApolloGraphQLInteraction()
      .uponReceiving("a hello request")
      .withQuery(query)
      .withRequest({
        path: "/graphql",
        method: "POST",
      })
      .willRespondWith({
        status: 200,
        headers: {
          "Content-Type": "application/json; charset=utf-8",
        },
        body: {
          data: {},
        },
      })
    return provider.addInteraction(graphqlQuery)
  })

  it("returns the correct response", () => {
    return expect(client.query({ query })).toEqual({ hello: "Hello world!" })
  })

  // verify with Pact, and reset expectations
  afterEach(() => provider.verify())
})