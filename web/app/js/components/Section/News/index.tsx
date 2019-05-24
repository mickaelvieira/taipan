import React from "react";
import { withStyles, WithStyles, createStyles } from "@material-ui/core/styles";
import Item from "./Item";
import Loader from "../../ui/Loader";
import NewsQuery, { query, variables, Data } from "../../apollo/Query/News";
import { Document } from "../../../types/document";
import FeedContainer from "../../ui/Feed/Container";

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

function hasReceivedData(data: Data | undefined): [boolean, Document[]] {
  let hasResults = false;
  let results: Document[] = [];

  if (data && "News" in data && "results" in data.News) {
    results = data.News.results;
    if (results.length > 0) {
      hasResults = true;
    }
  }

  return [hasResults, results];
}

export default function News() {
  return (
    <NewsQuery query={query} variables={variables}>
      {({ data, loading, error, fetchMore, networkStatus }) => {
        const [hasResults, documents] = hasReceivedData(data);
        console.log(hasResults);
        console.log(documents);
        // console.log(networkStatus);
        // console.log(fetchMore);

        return (
          <>
            {loading && <Loader />}
            {!loading && hasResults && (
              <FeedContainer>
                {documents.map((document: Document, index) => (
                  <Item document={document} index={index} key={document.id} />
                ))}
              </FeedContainer>
            )}
          </>
        );
      }}
    </NewsQuery>
  );
}
