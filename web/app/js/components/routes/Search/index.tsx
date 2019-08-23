import React from "react";
import { RouteSearchProps } from "../../../types/routes";
import FeedPage from "../../ui/Feed/Page";
import useSearch from "../../../hooks/useSearch";
import Feed from "./Feed";
import { SearchType } from "../../../types/search";
import { queryDocuments, queryBookmarks } from "../../apollo/Query/Search";
import ListBookmarks from "./List/Bookmark";
import ListDocuments from "./List/Document";
import ScrollToTop from "../../ui/ScrollToTop";

export interface SearchProps {
  type: SearchType;
  terms: string[];
}

export default function Search(_: RouteSearchProps): JSX.Element {
  const [type, terms] = useSearch();

  return (
    <ScrollToTop>
      <FeedPage>
        {type === "bookmark" && (
          <Feed
            name="searchbookmarks"
            terms={terms}
            type={type}
            query={queryBookmarks}
            List={ListBookmarks}
          />
        )}
        {type === "document" && (
          <Feed
            name="searchnews"
            terms={terms}
            type={type}
            query={queryDocuments}
            List={ListDocuments}
          />
        )}
      </FeedPage>
    </ScrollToTop>
  );
}
