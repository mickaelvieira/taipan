import React from "react";
import Layout from "../../Layout";
import ScrollToTop from "../../ui/ScrollToTop";
import Feed from "../../ui/Feed/Feed";
import List from "./List";
import { queryFavorites } from "../../apollo/Query/Feed";
import { subscriptionFavorites } from "../../apollo/Subscription/Feed";

export default function Favorites() {
  return (
    <Layout>
      <ScrollToTop>
        <Feed
          List={List}
          query={queryFavorites}
          subscription={subscriptionFavorites}
        />
      </ScrollToTop>
    </Layout>
  );
}
