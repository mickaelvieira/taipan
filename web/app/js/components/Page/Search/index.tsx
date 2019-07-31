import React from "react";
import Layout from "../../Layout/Search";
import ScrollToTop from "../../ui/ScrollToTop";
import useSearch from "../../../hooks/useSearch";
import Bookmarks from "./Bookmarks";
import Documents from "./Documents";

export default function Search(): JSX.Element {
  const [type, terms] = useSearch();
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
