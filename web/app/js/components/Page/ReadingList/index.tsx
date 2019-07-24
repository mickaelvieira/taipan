import React from "react";
import Layout from "../../Layout/Feed";
import ScrollToTop from "../../ui/ScrollToTop";
import Feed from "../../ui/Feed/Feed";
import List from "./List";
import { queryReadingList } from "../../apollo/Query/Feed";

export default function ReadingList(): JSX.Element {
  return (
    <Layout>
      <ScrollToTop>
        <Feed List={List} query={queryReadingList} />
      </ScrollToTop>
    </Layout>
  );
}
