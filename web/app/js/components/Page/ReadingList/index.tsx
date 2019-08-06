import React from "react";
import { RouteFeedProps } from "../../../types/routes";
import FeedPage from "../../ui/Feed/Page";
import Feed from "../../ui/Feed/Feed";
import List from "./List";
import { queryReadingList } from "../../apollo/Query/Feed";

export default function ReadingList(_: RouteFeedProps): JSX.Element {
  return (
    <FeedPage>
      <Feed List={List} query={queryReadingList} />
    </FeedPage>
  );
}
