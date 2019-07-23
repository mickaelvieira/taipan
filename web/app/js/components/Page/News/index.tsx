import React from "react";
import Layout from "../../Layout/Feed";
import ScrollToTop from "../../ui/ScrollToTop";
import Feed from "../../ui/Feed/Feed";
import List from "./List";
import { queryNews } from "../../apollo/Query/Feed";

export default function News(): JSX.Element {
  return (
    <Layout>
      <ScrollToTop>
        <Feed List={List} query={queryNews} />
      </ScrollToTop>
    </Layout>
  );
}
