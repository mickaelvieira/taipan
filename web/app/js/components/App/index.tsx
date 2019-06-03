import React from "react";
import { ApolloProvider } from "react-apollo";
import { createMuiTheme, MuiThemeProvider } from "@material-ui/core/styles";
import CssBaseline from "@material-ui/core/CssBaseline";
import { BrowserRouter, Route, Switch } from "react-router-dom";
import getApolloClient from "../../services/apollo";
import { NewsSection, ReadingListSection, FavoritesSection } from "../Section";
import uiTheme from "../ui/theme";
import Loader from "../ui/Loader";
import UserQuery from "../apollo/Query/User";
import { UserContext } from "../context";

export default function App() {
  const client = getApolloClient();
  const theme = createMuiTheme(uiTheme);

  console.log(theme);

  return (
    <ApolloProvider client={client}>
      <CssBaseline />
      <MuiThemeProvider theme={theme}>
        <BrowserRouter>
          <UserQuery>
            {({ data, loading, error }) => {
              if (loading) {
                return <Loader />;
              }

              if (error) {
                return <p>Oops</p>;
              }

              return !data || !data.User ? null : (
                <UserContext.Provider value={data.User}>
                  <Switch>
                    <Route exact path="/" component={NewsSection} />
                    <Route
                      exact
                      path="/reading-list"
                      component={ReadingListSection}
                    />
                    <Route
                      exact
                      path="/favorites"
                      component={FavoritesSection}
                    />
                  </Switch>
                </UserContext.Provider>
              );
            }}
          </UserQuery>
        </BrowserRouter>
      </MuiThemeProvider>
    </ApolloProvider>
  );
}
