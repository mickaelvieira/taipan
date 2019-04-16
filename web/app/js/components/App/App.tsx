import React, { useState, useEffect, useCallback } from "react";
import { Dispatch } from "redux";
import { connect } from "react-redux";
import { BrowserRouter as Router, Route } from "react-router-dom";
import Header from "components/Layout/Header";
import Grid from "components/Layout/Grid";
import Home from "components/Home";
import Feed from "components/Feed";
import Loader from "components/ui/Loader";
import AddBookmark from "components/Panels/AddBookmark";
import { fetchItems } from "store/actions/feed";
import { fetchUser } from "store/actions/user";
import { fetchIndex } from "store/actions/index";

interface Props {
  fetchIndex: () => Promise<any>;
  fetchItems: () => Promise<any>;
  fetchUser: () => Promise<any>;
}

function App({ fetchIndex, fetchItems, fetchUser }: Props) {
  const [isLoading, setIsLoading] = useState(true);
  const [isPanelOpen, setIsPanelOpen] = useState(false);

  useEffect(() => {
    (async () => {
      await fetchIndex();
      return Promise.all([fetchUser(), fetchItems()]);
    })()
      .then(() => setIsLoading(false))
      .catch(() => setIsLoading(false));
  }, []);

  const openAddBookmarkPanel = useCallback(() => setIsPanelOpen(true), []);
  const closeAddBookmarkPanel = useCallback(() => setIsPanelOpen(false), []);

  return (
    <Router>
      <Grid>
        <Header onClickAddBookmark={openAddBookmarkPanel} />
        {isLoading ? (
          <Loader />
        ) : (
          <>
            <Route
              exact
              path="/"
              render={props => (
                <Home {...props} onClickAddBookmark={openAddBookmarkPanel} />
              )}
            />
            <Route exact path="/feed" component={Feed} />
          </>
        )}
        <AddBookmark
          isOpen={isPanelOpen}
          onClickClose={closeAddBookmarkPanel}
        />
      </Grid>
    </Router>
  );
}

const mapDispatchToProps = (dispatch: Dispatch) => ({
  fetchIndex: () => dispatch(fetchIndex()),
  fetchUser: () => dispatch(fetchUser()),
  fetchItems: () => dispatch(fetchItems())
});

export default connect(
  undefined,
  mapDispatchToProps
)(App);
