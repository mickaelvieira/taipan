import React, { useState, useEffect, useCallback } from "react";
import { ApolloProvider, Query } from "react-apollo";
import { Dispatch } from "redux";
import { connect } from "react-redux";
import { BrowserRouter as Router, Route } from "react-router-dom";
import Header from "../Layout/Header";
import Grid from "../Layout/Grid";
import Home from "../Home";
import Feed from "../Feed";
import Loader from "../ui/Loader";
import AddBookmark from "../Panels/AddBookmark";
import { fetchItems } from "../../store/actions/feed";
import { fetchUser } from "../../store/actions/user";
import { fetchIndex } from "../../store/actions/index";
import getApolloClient from "../../services/apollo";

import query from "../../services/apollo/query/latest.graphql";

import Layout from "../Layout";

interface Props {
  fetchIndex: () => Promise<any>;
  fetchItems: () => Promise<any>;
  fetchUser: () => Promise<any>;
}

function App({ fetchIndex, fetchItems, fetchUser }: Props) {
  const [isLoading, setIsLoading] = useState(true);
  const [isPanelOpen, setIsPanelOpen] = useState(false);

  const client = getApolloClient();

  // useEffect(() => {
  //   (async () => {
  //     await fetchIndex();
  //     return Promise.all([fetchUser(), fetchItems()]);
  //   })()
  //     .then(() => setIsLoading(false))
  //     .catch(() => setIsLoading(false));
  // }, []);

  // const openAddBookmarkPanel = useCallback(() => setIsPanelOpen(true), []);
  // const closeAddBookmarkPanel = useCallback(() => setIsPanelOpen(false), []);

  return (
    <ApolloProvider client={client}>
      <Router>
        <Query query={query}>
          {({ data, loading, error }) => {
            console.log(data);

            if (loading) {
              return <Loader />;
            }

            return (
              <Layout>
                <div>Hello</div>
              </Layout>
            );
          }}
        </Query>
      </Router>
    </ApolloProvider>
  );
}

export default App;

// <Grid>
//   <Header onClickAddBookmark={openAddBookmarkPanel} />
//   {isLoading ? (
//     <Loader />
//   ) : (
//       <>
//         <Route
//           exact
//           path="/"
//           render={props => (
//             <Home {...props} onClickAddBookmark={openAddBookmarkPanel} />
//           )}
//         />
//         <Route exact path="/feed" component={Feed} />
//       </>
//     )}
//   <AddBookmark
//     isOpen={isPanelOpen}
//     onClickClose={closeAddBookmarkPanel}
//   />
// </Grid>

// const mapDispatchToProps = (dispatch: Dispatch) => ({
//   fetchIndex: () => dispatch(fetchIndex()),
//   fetchUser: () => dispatch(fetchUser()),
//   fetchItems: () => dispatch(fetchItems())
// });

// export default connect(
//   undefined,
//   mapDispatchToProps
// )(App);
