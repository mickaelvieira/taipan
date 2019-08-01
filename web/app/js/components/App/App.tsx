import React from "react";
import { ApolloProvider } from "react-apollo";
import { createMuiTheme, MuiThemeProvider } from "@material-ui/core/styles";
import CssBaseline from "@material-ui/core/CssBaseline";
import { BrowserRouter, Route, Switch } from "react-router-dom";
import getApolloClient, { genRandomId } from "../../services/apollo";
import Layout from "../Layout/Layout";
import {
  ErrorPage,
  NewsPage,
  ReadingListPage,
  FavoritesPage,
  SearchPage,
  SyndicationPage,
  AccountPage
} from "../Page";
import getThemeOptions from "../ui/themes";
import Loader from "../ui/Loader";
import InitQuery, { Data } from "../apollo/Query/Init";
import { ClientContext, AppContext } from "../context";
import AppUser from "./User";
import AppFeeds from "./Feeds";

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
  const clientId = genRandomId();
  const client = getApolloClient(clientId);

  return (
    <ClientContext.Provider value={clientId}>
      <ApolloProvider client={client}>
        <CssBaseline />
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
              const { loggedIn: user } = users;
              const theme = createMuiTheme(getThemeOptions(user.theme));
              console.log(theme);

              return (
                <MuiThemeProvider theme={theme}>
                  <AppContext.Provider value={app.info}>
                    <AppUser loggedIn={user}>
                      <AppFeeds client={client}>
                        <Layout>
                          {props => (
                            <Switch>
                              <Route
                                exact
                                path="/"
                                render={routeProps => (
                                  <NewsPage {...routeProps} {...props} />
                                )}
                              />
                              <Route
                                exact
                                path="/reading-list"
                                render={routeProps => (
                                  <ReadingListPage {...routeProps} {...props} />
                                )}
                              />
                              <Route
                                exact
                                path="/favorites"
                                render={routeProps => (
                                  <FavoritesPage {...routeProps} {...props} />
                                )}
                              />
                              <Route
                                exact
                                path="/search/:type?"
                                render={routeProps => (
                                  <SearchPage {...routeProps} {...props} />
                                )}
                              />
                              <Route
                                exact
                                path="/syndication"
                                render={routeProps => (
                                  <SyndicationPage {...routeProps} {...props} />
                                )}
                              />
                              <Route
                                exact
                                path="/account"
                                render={routeProps => (
                                  <AccountPage {...routeProps} {...props} />
                                )}
                              />
                            </Switch>
                          )}
                        </Layout>
                      </AppFeeds>
                    </AppUser>
                  </AppContext.Provider>
                </MuiThemeProvider>
              );
            }}
          </InitQuery>
        </BrowserRouter>
      </ApolloProvider>
    </ClientContext.Provider>
  );
}
