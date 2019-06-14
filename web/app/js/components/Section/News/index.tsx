import React, { useState, useRef, useEffect } from "react";
import Layout from "../../Layout";
import ScrollToTop from "../../ui/ScrollToTop";
import Feed from "../../ui/Feed/Feed";
import List from "./List";
import { queryNews } from "../../apollo/Query/Feed";
import LatestNewsQuery from "../../apollo/Query/LatestNews";
import { NewsContext } from "../../context";

export default function News() {
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
          <LatestNewsQuery variables={{ pagination: { limit: 10, to: toId } }}>
            {({ data, client, startPolling }) => {
              console.log("data");
              console.log(data);
              poll.current = startPolling;

              return !data || !data.News ? null : (
                <div>See {data.News.results.length} news documents</div>
              );
            }}
          </LatestNewsQuery>
          <Feed List={List} query={queryNews} />
        </NewsContext.Provider>
      </ScrollToTop>
    </Layout>
  );
}
