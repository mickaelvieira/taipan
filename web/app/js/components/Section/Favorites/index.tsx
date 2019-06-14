import React from "react";
import Layout from "../../Layout";
import ScrollToTop from "../../ui/ScrollToTop";
import Feed from "../../ui/Feed/Feed";
import List from "./List";
import { queryFavorites } from "../../apollo/Query/Feed";
import LatestFavoriteSubscription from "../../apollo/Subscription/Favorite";

export default function Favorites() {
  return (
    <Layout>
      <ScrollToTop>
        <LatestFavoriteSubscription>{() => <></>}</LatestFavoriteSubscription>
        <Feed List={List} query={queryFavorites} />
      </ScrollToTop>
    </Layout>
  );
}
