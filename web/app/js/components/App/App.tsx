import React from "react";
import { ApolloProvider } from "react-apollo";
import { createMuiTheme, MuiThemeProvider } from "@material-ui/core/styles";
import CssBaseline from "@material-ui/core/CssBaseline";
import { BrowserRouter, Route, Switch } from "react-router-dom";
import getApolloClient, { genRandomId } from "../../services/apollo";
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
                            path="/search/:type?"
                            component={SearchPage}
                          />
                          <Route
                            exact
                            path="/syndication"
                            component={SyndicationPage}
                          />
                          <Route
                            exact
                            path="/account"
                            component={AccountPage}
                          />
                        </Switch>
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
