import React from "react";
import Paper from "@material-ui/core/Paper";
import Layout from "../../Layout/Subscriptions";
import ScrollToTop from "../../ui/ScrollToTop";
import Table from "./Table";

export default function Subscriptions(): JSX.Element {
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
