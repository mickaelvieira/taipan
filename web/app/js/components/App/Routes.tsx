import React from "react";
import { Route, Switch } from "react-router-dom";
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

export default function Routes(): JSX.Element {
  return (
    <Switch>
      <Route
        exact
        path="/"
        render={routeProps => <NewsPage {...routeProps} />}
      />
      <Route
        exact
        path="/reading-list"
        render={routeProps => <ReadingListPage {...routeProps} />}
      />
      <Route
        exact
        path="/favorites"
        render={routeProps => <FavoritesPage {...routeProps} />}
      />
      <Route
        exact
        path="/search/:type?"
        render={routeProps => <SearchPage {...routeProps} />}
      />
      <Route
        exact
        path="/syndication"
        render={routeProps => <SyndicationPage {...routeProps} />}
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
  );
}
