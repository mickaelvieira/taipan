import React, { useState, useRef, useEffect } from "react";
import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import Layout from "../../Layout";
import ScrollToTop from "../../ui/ScrollToTop";
import Feed from "../../ui/Feed/Feed";
import List from "./List";
import { queryNews } from "../../apollo/Query/Feed";
import { addItemsFromFeedResults } from "../../apollo/helpers/feed";
import LatestNewsQuery, {
  query as queryLatestNews
} from "../../apollo/Query/LatestNews";
import { NewsContext } from "../../context";
import { subscriptionNews } from "../../apollo/Subscription/Feed";

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

  useEffect(() => {
    if (toId && poll.current) {
      poll.current(30000);
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
              poll.current = startPolling;

              return !data ||
                !data.News ||
                data.News.results.length === 0 ? null : (
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
          <Feed List={List} query={queryNews} subscription={subscriptionNews} />
        </NewsContext.Provider>
      </ScrollToTop>
    </Layout>
  );
}
