import React from "react";
import Layout from "../../Layout/Account";
import ScrollToTop from "../../ui/ScrollToTop";
import Prolile from "./Profile";
import { PageTitle } from "../../ui/Title";

export default function Account(): JSX.Element | null {
  return (
    <Layout>
      <ScrollToTop>
        <PageTitle value="My Account" />
        <Prolile />
      </ScrollToTop>
    </Layout>
  );
}
