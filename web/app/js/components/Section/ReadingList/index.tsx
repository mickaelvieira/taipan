import React from "react";
import Item from "./Item";
import Loader from "../../ui/Loader";
import ReadingListQuery, {
  query,
  variables,
  Data
} from "../../apollo/Query/ReadingList";
import { Bookmark } from "../../../types/bookmark";
import FeedContainer from "../../ui/Feed/Container";

function hasReceivedData(data: Data | undefined): [boolean, Bookmark[]] {
  let hasResults = false;
  let results: Bookmark[] = [];

  if (data && "GetReadingList" in data && "results" in data.GetReadingList) {
    results = data.GetReadingList.results;
    if (results.length > 0) {
      hasResults = true;
    }
  }

  return [hasResults, results];
}

// @TODO Add infinite scroll to all feeds

export default function News() {
  return (
    <ReadingListQuery query={query} variables={variables}>
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
    </ReadingListQuery>
  );
}
