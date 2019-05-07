import React, { useState } from "react";
import { withStyles, WithStyles, createStyles } from "@material-ui/core/styles";
import SwipeableViews from 'react-swipeable-views';
import AppBar from '@material-ui/core/AppBar';
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
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
    },
    tabs: {
      width: "100%"
    }
  });

interface Props extends WithStyles<typeof styles> {
  bookmark: UserBookmark;
}

function hasReceivedBookmarks(
  data: Data | undefined,
  queyrKey: string
): [boolean, UserBookmark[]] {
  let hasResults = false;
  let results: UserBookmark[] = [];

  if (
    data &&
    queyrKey in data &&
    "results" in data[queyrKey]
  ) {
    results = data[queyrKey].results;
    if (results.length > 0) {
      hasResults = true;
    }
  }

  return [hasResults, results];
}

export default withStyles(styles)(function Feed({ classes }: Props) {
  const [tabIndex, setTabIndex] = useState(0)

  console.log(tabIndex)

  return (
    <>
      <Tabs
        value={tabIndex}
        onChange={(event, index) => setTabIndex(index)}
        indicatorColor="primary"
        textColor="primary"
        variant="fullWidth"
        className={classes.tabs}
      >
        <Tab label="Latest" />
        <Tab label="Bookmarks" />
      </Tabs>
      <LatestBookmarksQuery query={query} variables={variables}>
        {({ data, loading, error, fetchMore, networkStatus }) => {
          const [hasLatest, latest] = hasReceivedBookmarks(data, "GetLatestBookmarks");
          const [hasNewest, newest] = hasReceivedBookmarks(data, "GetNewBookmarks");
          console.log(hasLatest);
          console.log(latest);
          console.log(networkStatus);
          console.log(fetchMore);

          return (
            <SwipeableViews
              index={tabIndex}
              onChangeIndex={setTabIndex}
            >
              <FeedWrapper
                queryKey="GetNewBookmarks"
                isLoading={loading}
                fetchMore={fetchMore}
                hasResults={hasNewest}
                bookmarks={newest}
              />
              <FeedWrapper
                queryKey="GetLatestBookmarks"
                isLoading={loading}
                fetchMore={fetchMore}
                hasResults={hasLatest}
                bookmarks={latest}
              />
            </SwipeableViews>
          );
        }}
      </LatestBookmarksQuery>
    </>
  );
});
