import React from "react";
import Layout from "../../Layout/Search";
import ScrollToTop from "../../ui/ScrollToTop";

import { RouteSearchProps } from "../../../types/routes";
// interface Props extends RouteSearchProps {}

export default function Search({
  match: { params }
}: RouteSearchProps): JSX.Element {
  console.log(params);

  return (
    <Layout>
      <ScrollToTop>
        <div>search</div>
      </ScrollToTop>
    </Layout>
  );
}
