import React from "react";
import Layout from "../../Layout";
import ScrollToTop from "../../ui/ScrollToTop";
import Feed from "./Feed";

export default function Favorites() {
  return (
    <Layout>
      <ScrollToTop>
        <Feed />
      </ScrollToTop>
    </Layout>
  );
}
