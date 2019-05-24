import React from "react";
import Item from "./Item";
import Loader from "../../ui/Loader";
import FavoritesQuery, {
  query,
  variables,
  Data
} from "../../apollo/Query/Favorites";
import { Bookmark } from "../../../types/bookmark";
import FeedContainer from "../../ui/Feed/Container";

function hasReceivedData(data: Data | undefined): [boolean, Bookmark[]] {
  let hasResults = false;
  let results: Bookmark[] = [];

  if (data && "GetFavorites" in data && "results" in data.GetFavorites) {
    results = data.GetFavorites.results;
    if (results.length > 0) {
      hasResults = true;
    }
  }

  return [hasResults, results];
}

export default function News() {
  return (
    <FavoritesQuery query={query} variables={variables}>
      {({ data, loading, error, fetchMore, networkStatus }) => {
        const [hasResults, bookmarks] = hasReceivedData(data);
        console.log(hasResults);
        console.log(bookmarks);
        // console.log(networkStatus);
        // console.log(fetchMore);

        return (
          <>
            {loading && <Loader />}
            {!loading && hasResults && (
              <FeedContainer>
                {bookmarks.map((bookmark: Bookmark, index) => (
                  <Item bookmark={bookmark} index={index} key={bookmark.id} />
                ))}
              </FeedContainer>
            )}
          </>
        );
      }}
    </FavoritesQuery>
  );
}
