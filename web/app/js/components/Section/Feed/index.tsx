import React from "react";
import { withStyles, WithStyles, createStyles } from "@material-ui/core/styles";
import Loader from "../../ui/Loader";
import PointerEvents from "../../ui/PointerEvents";
import FeedItem from "./Item";
import FeedWrapper from "./Wrapper";
import List from "./List";

import LatestBookmarksQuery, {
  query,
  variables,
  Data
} from "../../apollo/Query/LatestBookmarks";
import { UserBookmark } from "../../../types/bookmark";

const styles = () =>
  createStyles({
    container: {
      display: "flex",
      flexDirection: "column"
    }
  });

interface Props extends WithStyles<typeof styles> {
  bookmark: UserBookmark;
}

function hasReceivedBookmarks(
  data: Data | undefined
): [boolean, UserBookmark[]] {
  let hasResults = false;
  let results: UserBookmark[] = [];

  if (
    data &&
    "GetLatestBookmarks" in data &&
    "results" in data.GetLatestBookmarks
  ) {
    results = data.GetLatestBookmarks.results;
    if (results.length > 0) {
      hasResults = true;
    }
  }

  return [hasResults, results];
}

export default withStyles(styles)(function Feed({ classes }: Props) {
  return (
    <LatestBookmarksQuery query={query} variables={variables}>
      {({ data, loading, error, fetchMore, networkStatus }) => {
        const [hasResults, bookmarks] = hasReceivedBookmarks(data);
        console.log(hasResults);
        console.log(bookmarks);
        console.log(networkStatus);
        console.log(fetchMore);

        return (
          <FeedWrapper
            isLoading={loading}
            fetchMore={fetchMore}
            hasResults={hasResults}
            bookmarks={bookmarks}
          />
        );
      }}
    </LatestBookmarksQuery>
  );
});
