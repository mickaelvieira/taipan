import React from "react";
import Layout from "../../Layout/Syndication";
import ScrollToTop from "../../ui/ScrollToTop";
import { PageTitle } from "../../ui/Title";
import Table from "./Table";

export default function Syndication(): JSX.Element {
  return (
    <Layout>
      <ScrollToTop>
        <PageTitle value="Web syndication" />
        <Table />
      </ScrollToTop>
    </Layout>
  );
}
