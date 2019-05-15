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

import FeedsQuery, {
  query,
  variables,
  Data
} from "../../apollo/Query/Feeds";
import { Bookmark } from "../../../types/bookmark";

const styles = () =>
  createStyles({
    tabs: {
      width: "100%"
    }
  });

interface Props extends WithStyles<typeof styles> {
  bookmark: Bookmark;
}

function hasReceivedBookmarks(
  data: Data | undefined,
  queyrKey: string
): [boolean, Bookmark[]] {
  let hasResults = false;
  let results: Bookmark[] = [];

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
      <FeedsQuery query={query} variables={variables}>
        {({ data, loading, error, fetchMore, networkStatus }) => {
          const [hasLatest, latest] = hasReceivedBookmarks(data, "GetLatestBookmarks");
          const [hasNewest, newest] = hasReceivedBookmarks(data, "GetLatestDocuments");
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
                queryKey="GetLatestDocuments"
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
      </FeedsQuery>
    </>
  );
});
