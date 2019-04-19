import React from "react";
import { ApolloProvider } from "react-apollo";
import CssBaseline from "@material-ui/core/CssBaseline";
import { BrowserRouter, Route, Switch } from "react-router-dom";
import getApolloClient from "../../services/apollo";
import { HomeSection, FeedSection } from "../Section";

import Layout from "../Layout";

export default function App() {
  const client = getApolloClient();

  return (
    <ApolloProvider client={client}>
      <CssBaseline />
      <BrowserRouter>
        <Layout>
          <Switch>
            <Route exact path="/" component={HomeSection} />
            <Route exact path="/feed" component={FeedSection} />
          </Switch>
        </Layout>
      </BrowserRouter>
    </ApolloProvider>
  );
}
