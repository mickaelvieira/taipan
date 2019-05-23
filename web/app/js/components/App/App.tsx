import React from "react";
import { ApolloProvider } from "react-apollo";
import { createMuiTheme, MuiThemeProvider } from "@material-ui/core/styles";
import CssBaseline from "@material-ui/core/CssBaseline";
import { BrowserRouter, Route, Switch } from "react-router-dom";
import getApolloClient from "../../services/apollo";
import { NewsSection, ReadingListSection, FavoriteSection } from "../Section";
import uiTheme from "../ui/theme";
import Layout from "../Layout";

export default function App() {
  const client = getApolloClient();
  const theme = createMuiTheme(uiTheme);

  console.log(theme);

  return (
    <ApolloProvider client={client}>
      <CssBaseline />
      <MuiThemeProvider theme={theme}>
        <BrowserRouter>
          <Layout>
            <Switch>
              <Route exact path="/" component={NewsSection} />
              <Route
                exact
                path="/reading-list"
                component={ReadingListSection}
              />
              <Route exact path="/favorites" component={FavoriteSection} />
            </Switch>
          </Layout>
        </BrowserRouter>
      </MuiThemeProvider>
    </ApolloProvider>
  );
}
