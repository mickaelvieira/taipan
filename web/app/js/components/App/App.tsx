import React from "react";
import { ApolloProvider } from "react-apollo";
import { createMuiTheme, MuiThemeProvider } from "@material-ui/core/styles";
import CssBaseline from "@material-ui/core/CssBaseline";
import { BrowserRouter, Route, Switch } from "react-router-dom";
import getApolloClient from "../../services/apollo";
import { NewsSection, ReadingListSection, FavoritesSection } from "../Section";
import uiTheme from "../ui/theme";

export default function App() {
  const client = getApolloClient();
  const theme = createMuiTheme(uiTheme);

  console.log(theme);

  return (
    <ApolloProvider client={client}>
      <CssBaseline />
      <MuiThemeProvider theme={theme}>
        <BrowserRouter>
          <Switch>
            <Route exact path="/" component={NewsSection} />
            <Route exact path="/reading-list" component={ReadingListSection} />
            <Route exact path="/favorites" component={FavoritesSection} />
          </Switch>
        </BrowserRouter>
      </MuiThemeProvider>
    </ApolloProvider>
  );
}
