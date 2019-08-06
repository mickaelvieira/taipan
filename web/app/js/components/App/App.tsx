import React from "react";
import { ApolloProvider } from "react-apollo";
import { createMuiTheme, MuiThemeProvider } from "@material-ui/core/styles";
import CssBaseline from "@material-ui/core/CssBaseline";
import { BrowserRouter, Route, Switch } from "react-router-dom";
import getApolloClient, { genRandomId } from "../../services/apollo";
import Layout from "../Layout";
import {
  NotFoundPage,
  LoginPage,
  NewsPage,
  ReadingListPage,
  FavoritesPage,
  SearchPage,
  SyndicationPage,
  AccountPage
} from "../Page";
import getThemeOptions from "../ui/themes";
import Loader from "../ui/Loader";
import InitQuery from "../apollo/Query/Init";
import { ClientContext, AppContext } from "../context";

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

              let user = null;
              let appInfo = { info: { name: "Taipan", version: "0.0.0" } };
              if (data && !error) {
                let { users, app } = data;
                if (users.loggedIn) {
                  user = users.loggedIn;
                }
                if (app) {
                  appInfo = app;
                }
              }

              const theme = createMuiTheme(
                getThemeOptions(user ? user.theme : null)
              );

              console.log(theme);

              return (
                <MuiThemeProvider theme={theme}>
                  <AppContext.Provider value={appInfo.info}>
                    <Layout user={user}>
                      <Switch>
                        <Route
                          exact
                          path="/"
                          render={routeProps => <NewsPage {...routeProps} />}
                        />
                        <Route
                          exact
                          path="/reading-list"
                          render={routeProps => (
                            <ReadingListPage {...routeProps} />
                          )}
                        />
                        <Route
                          exact
                          path="/favorites"
                          render={routeProps => (
                            <FavoritesPage {...routeProps} />
                          )}
                        />
                        <Route
                          exact
                          path="/search/:type?"
                          render={routeProps => <SearchPage {...routeProps} />}
                        />
                        <Route
                          exact
                          path="/syndication"
                          render={routeProps => (
                            <SyndicationPage {...routeProps} />
                          )}
                        />
                        <Route
                          exact
                          path="/account"
                          render={routeProps => <AccountPage {...routeProps} />}
                        />
                        <Route
                          exact
                          path="/sign-in"
                          render={routeProps => <LoginPage {...routeProps} />}
                        />
                        <Route component={NotFoundPage} />
                      </Switch>
                    </Layout>
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
