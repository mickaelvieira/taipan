import React, { useState } from "react";
import { makeStyles } from "@material-ui/core/styles";
import SwipeableViews from "react-swipeable-views";
import AppBar from "@material-ui/core/AppBar";
import Tabs from "@material-ui/core/Tabs";
import Tab from "@material-ui/core/Tab";
import Loader from "../Loader";
import PointerEvents from "../PointerEvents";
import FeedItem from "./Item";
import FeedWrapper from "./Wrapper";
import List from "./List";

import FeedsQuery, {
  query,
  variables,
  Data
} from "../../apollo/Query/ReadingList";
import { Bookmark } from "../../../types/bookmark";

const useStyles = makeStyles({
  tabs: {
    width: "100%"
  },
  container: {
    width: "100%",
    minHeight: 200
  }
});

interface Props {
  bookmark: Bookmark;
}

function hasReceivedBookmarks(
  data: Data | undefined,
  queyrKey: string
): [boolean, Bookmark[]] {
  let hasResults = false;
  let results: Bookmark[] = [];

  if (data && queyrKey in data && "results" in data[queyrKey]) {
    results = data[queyrKey].results;
    if (results.length > 0) {
      hasResults = true;
    }
  }

  return [hasResults, results];
}

export default function Feed() {
  const classes = useStyles()
  const [tabIndex, setTabIndex] = useState(0);

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
        <Tab label="News" />
        <Tab label="Reading List" />
        <Tab label="Bookmarks" />
      </Tabs>
      <FeedsQuery query={query} variables={variables}>
        {({ data, loading, error, fetchMore, networkStatus }) => {
          const [hasLatest, latest] = hasReceivedBookmarks(
            data,
            "GetLatestBookmarks"
          );
          const [hasNewest, newest] = hasReceivedBookmarks(
            data,
            "GetLatestDocuments"
          );
          // console.log(hasLatest);
          // console.log(latest);
          // console.log(networkStatus);
          // console.log(fetchMore);

          return (
            <SwipeableViews
              index={tabIndex}
              disabled
              animateHeight
              onChangeIndex={setTabIndex}
              className={classes.container}
            >
              <FeedWrapper
                queryKey="GetLatestDocuments"
                isLoading={loading}
                fetchMore={fetchMore}
                hasResults={hasNewest}
                bookmarks={newest}
              />
              <div>Reading list</div>
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
