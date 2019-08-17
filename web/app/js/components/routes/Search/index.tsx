import React from "react";
import { RouteSearchProps } from "../../../types/routes";
import Grid from "../../ui/Grid";
import useSearch from "../../../hooks/useSearch";
import Bookmarks from "./Bookmarks";
import Documents from "./Documents";

export default function Search(_: RouteSearchProps): JSX.Element {
  const [type, terms] = useSearch();
  return (
    <Grid>
      {type === "bookmark" && <Bookmarks terms={terms} type={type} />}
      {type === "document" && <Documents terms={terms} type={type} />}
    </Grid>
  );
}
