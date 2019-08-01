import React from "react";
import { LayoutRenderProps } from "../../Layout/Layout";
import FeedPage from "../../ui/Feed/Page";
import Feed from "../../ui/Feed/Feed";
import List from "./List";
import { queryNews } from "../../apollo/Query/Feed";

export default function News(props: LayoutRenderProps): JSX.Element {
  return (
    <FeedPage {...props}>
      <Feed List={List} query={queryNews} />
    </FeedPage>
  );
}
