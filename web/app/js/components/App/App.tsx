import React from "react";
import { ApolloProvider } from "@apollo/react-hooks";
import { BrowserRouter } from "react-router-dom";
import getApolloClient, { genRandomId } from "../../services/apollo";
import { ClientContext } from "../context";
import Environment from "./Environment";

export default function App(): JSX.Element {
  const clientId = genRandomId();
  const client = getApolloClient(clientId);

  // @TODO move the BrowserRouter to the Routes component, Layout does not necessarly need to know about the router
  return (
    <ClientContext.Provider value={clientId}>
      <ApolloProvider client={client}>
        <BrowserRouter>
          <Environment />
        </BrowserRouter>
      </ApolloProvider>
    </ClientContext.Provider>
  );
}
