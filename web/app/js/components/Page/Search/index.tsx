import React from "react";
import Grid from "../../ui/Grid";
import useSearch from "../../../hooks/useSearch";
import Bookmarks from "./Bookmarks";
import Documents from "./Documents";

export default function Search(): JSX.Element {
  const [type, terms] = useSearch();
  return (
    <Grid>
      {type === "bookmark" && <Bookmarks terms={terms} type={type} />}
      {type === "document" && <Documents terms={terms} type={type} />}
    </Grid>
  );
}
