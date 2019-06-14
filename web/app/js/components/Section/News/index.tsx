import React, { useState, useRef, useEffect } from "react";
import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import Layout from "../../Layout";
import ScrollToTop from "../../ui/ScrollToTop";
import Feed from "../../ui/Feed/Feed";
import List from "./List";
import { queryNews, addItemsFromFeedResults } from "../../apollo/Query/Feed";
import LatestNewsQuery, {
  query as queryLatestNews
} from "../../apollo/Query/LatestNews";
import { NewsContext } from "../../context";

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    button: {
      margin: theme.spacing(1)
    }
  })
);

const mergeResults = client => {
  try {
    const oldData = client.readQuery({ query: queryNews });
    const newData = client.readQuery({ query: queryLatestNews });
    console.log(oldData);
    console.log(newData);
    if (oldData && newData) {
      const [newsResult, latestNewsResults] = addItemsFromFeedResults(
        oldData.News,
        newData.News
      );
      client.writeQuery({
        query: queryNews,
        data: { News: newsResult }
      });

      client.writeQuery({
        query: queryLatestNews,
        data: { News: latestNewsResults }
      });
    }
  } catch (e) {
    console.warn(e);
  }
};

export default function News() {
  const classes = useStyles();
  const [toId, setToId] = useState("");
  const poll = useRef<(time: number) => void | undefined>();

  // console.log("toId");
  // console.log(toId);

  useEffect(() => {
    if (toId && poll.current) {
      console.log("start polling");
      console.log("toId");
      console.log(toId);
      poll.current(5000);
    }
  }, [toId]);

  return (
    <Layout>
      <ScrollToTop>
        <NewsContext.Provider value={setToId}>
          <LatestNewsQuery
            skip={toId == ""}
            variables={{ pagination: { limit: 10, to: toId } }}
          >
            {({ data, client, startPolling, stopPolling }) => {
              console.log("data");
              console.log(data);
              poll.current = startPolling;

              return !data || !data.News ? null : (
                <div>
                  <Button
                    className={classes.button}
                    onClick={() => {
                      stopPolling();
                      mergeResults(client);
                    }}
                  >
                    See {data.News.results.length} latest news
                  </Button>
                </div>
              );
            }}
          </LatestNewsQuery>
          <Feed List={List} query={queryNews} />
        </NewsContext.Provider>
      </ScrollToTop>
    </Layout>
  );
}
