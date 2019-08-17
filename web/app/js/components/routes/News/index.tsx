import React from "react";
import { RouteFeedProps } from "../../../types/routes";
import FeedPage from "../../ui/Feed/Page";
import Feed from "../../ui/Feed/Feed";
import List from "./List";
import { queryNews } from "../../apollo/Query/Feed";

export default function News(_: RouteFeedProps): JSX.Element {
  return (
    <FeedPage>
      <Feed List={List} query={queryNews} name="news" />
    </FeedPage>
  );
}
