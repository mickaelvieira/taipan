import React from "react";
import Layout from "../../Layout/Account";
import ScrollToTop from "../../ui/ScrollToTop";
import Prolile from "./Profile";

export default function Account(): JSX.Element | null {
  return (
    <Layout>
      <ScrollToTop>
        <Prolile />
      </ScrollToTop>
    </Layout>
  );
}
