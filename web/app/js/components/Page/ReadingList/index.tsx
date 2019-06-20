import React from "react";
import Layout from "../../Layout";
import ScrollToTop from "../../ui/ScrollToTop";
import Feed from "../../ui/Feed/Feed";
import List from "./List";
import { queryReadingList } from "../../apollo/Query/Feed";
import { subscriptionReadingList } from "../../apollo/Subscription/Feed";

export default function ReadingList() {
  return (
    <Layout>
      <ScrollToTop>
        <Feed
          List={List}
          query={queryReadingList}
          subscription={subscriptionReadingList}
        />
      </ScrollToTop>
    </Layout>
  );
}
