import React, { lazy, Suspense } from "react";
import { Route, Switch } from "react-router-dom";
import Loader from "../ui/Loader";

const SignupPage = lazy(() =>
  import(/* webpackChunkName: "Signup" */ "./Signup")
);
const SigninPage = lazy(() =>
  import(/* webpackChunkName: "Signin" */ "./Signin")
);
const NotFoundPage = lazy(() =>
  import(/* webpackChunkName: "NotFound" */ "./NotFound")
);
const NewsPage = lazy(() => import(/* webpackChunkName: "News" */ "./News"));
const ReadingListPage = lazy(() =>
  import(/* webpackChunkName: "ReadingList" */ "./ReadingList")
);
const FavoritesPage = lazy(() =>
  import(/* webpackChunkName: "Favorites" */ "./Favorites")
);
const SearchPage = lazy(() =>
  import(/* webpackChunkName: "Search" */ "./Search")
);
const SubscriptionsPage = lazy(() =>
  import(/* webpackChunkName: "Subscriptions" */ "./Subscriptions")
);
const AccountPage = lazy(() =>
  import(/* webpackChunkName: "Account" */ "./Account")
);

export default function Routes(): JSX.Element {
  return (
    <Suspense fallback={<Loader />}>
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
          render={routeProps => <SubscriptionsPage {...routeProps} />}
        />
        <Route
          exact
          path="/account"
          render={routeProps => <AccountPage {...routeProps} />}
        />
        <Route
          exact
          path="/sign-in"
          render={routeProps => <SigninPage {...routeProps} />}
        />
        <Route
          exact
          path="/join"
          render={routeProps => <SignupPage {...routeProps} />}
        />
        <Route component={NotFoundPage} />
      </Switch>
    </Suspense>
  );
}