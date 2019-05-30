import React from "react";
import Layout from "../../Layout";
import ScrollToTop from "../../ui/ScrollToTop";
import Feed from "../../ui/Feed/Feed";
import List from "./List";
import { queryNews, DataKey } from "../../apollo/Query/Feed";

export default function News() {
  return (
    <Layout>
      <ScrollToTop>
        <Feed List={List} dataKey={DataKey.NEWS} query={queryNews} />
      </ScrollToTop>
    </Layout>
  );
}
