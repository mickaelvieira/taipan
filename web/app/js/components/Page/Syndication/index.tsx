import React from "react";
import Paper from "@material-ui/core/Paper";
import Layout from "../../Layout/Syndication";
import ScrollToTop from "../../ui/ScrollToTop";
import Table from "./Table";

export default function Syndication(): JSX.Element {
  return (
    <Layout>
      <ScrollToTop>
        <Paper>
          <Table />
        </Paper>
      </ScrollToTop>
    </Layout>
  );
}
