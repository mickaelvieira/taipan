import React from "react";
import { ApolloProvider } from "react-apollo";
import { createMuiTheme, MuiThemeProvider } from "@material-ui/core/styles";
import CssBaseline from "@material-ui/core/CssBaseline";
import { BrowserRouter, Route, Switch } from "react-router-dom";
import getApolloClient from "../../services/apollo";
import {
  ErrorPage,
  NewsPage,
  ReadingListPage,
  FavoritesPage,
  SyndicationPage
} from "../Page";
import uiTheme from "../ui/theme";
import Loader from "../ui/Loader";
import InitQuery, { Data } from "../apollo/Query/Init";
import { AppContext, UserContext } from "../context";

function canBoostrap(data: Data | undefined): boolean {
  if (!data) {
    return false;
  }
  if (!data.users) {
    return false;
  }
  if (!data.users.loggedIn) {
    return false;
  }
  if (!data.app) {
    return false;
  }
  if (!data.app.info) {
    return false;
  }
  return true;
}

export default function App(): JSX.Element {
  const client = getApolloClient();
  const theme = createMuiTheme(uiTheme);

  console.log(theme);

  return (
    <ApolloProvider client={client}>
      <CssBaseline />
      <MuiThemeProvider theme={theme}>
        <BrowserRouter>
          <InitQuery>
            {({ data, loading, error }) => {
              if (loading) {
                return <Loader />;
              }

              if (error) {
                return <ErrorPage error={error} />;
              }

              if (!canBoostrap(data)) {
                return null;
              }

              const { users, app } = data as Data;

              return (
                <AppContext.Provider value={app.info}>
                  <UserContext.Provider value={users.loggedIn}>
                    <Switch>
                      <Route exact path="/" component={NewsPage} />
                      <Route
                        exact
                        path="/reading-list"
                        component={ReadingListPage}
                      />
                      <Route
                        exact
                        path="/favorites"
                        component={FavoritesPage}
                      />
                      <Route
                        exact
                        path="/syndication"
                        component={SyndicationPage}
                      />
                    </Switch>
                  </UserContext.Provider>
                </AppContext.Provider>
              );
            }}
          </InitQuery>
        </BrowserRouter>
      </MuiThemeProvider>
    </ApolloProvider>
  );
}
