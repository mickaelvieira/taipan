import React from "react";
import Layout from "../../Layout/Search";
import ScrollToTop from "../../ui/ScrollToTop";
import { getSearchTerms, getSearchType } from "../../../helpers/search";
import Bookmarks from "./Bookmarks";
import Documents from "./Documents";

export default function Search(): JSX.Element {
  const url = new URL(`${document.location}`);
  const terms = getSearchTerms(url.searchParams.get("terms"));
  const type = getSearchType(url.pathname, url.searchParams.get("type"));

  console.log(type);
  console.log(terms);

  return (
    <Layout>
      <ScrollToTop>
        {type === "bookmark" && <Bookmarks terms={terms} />}
        {type === "document" && <Documents terms={terms} />}
      </ScrollToTop>
    </Layout>
  );
}
