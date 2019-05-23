import React from "react";
import { withStyles, WithStyles, createStyles } from "@material-ui/core/styles";
import Item from "./Item";
import Loader from "../../ui/Loader";
import FavoritesQuery, {
  query,
  variables,
  Data
} from "../../apollo/Query/Favorites";
import { Bookmark } from "../../../types/bookmark";

const styles = () =>
  createStyles({
    tabs: {
      width: "100%"
    },
    container: {
      width: "100%",
      minHeight: 200
    }
  });

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

export default withStyles(styles)(function News({
  classes
}: WithStyles<typeof styles>) {
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
              <div className={classes.container}>
                {bookmarks.map((bookmark: Bookmark, index) => (
                  <Item bookmark={bookmark} index={index} key={bookmark.id} />
                ))}
              </div>
            )}
          </>
        );
      }}
    </FavoritesQuery>
  );
});
