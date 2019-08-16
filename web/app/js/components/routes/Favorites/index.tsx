import React from "react";
import { RouteFeedProps } from "../../../types/routes";
import FeedPage from "../../ui/Feed/Page";
import Feed from "../../ui/Feed/Feed";
import List from "./List";
import { queryFavorites } from "../../apollo/Query/Feed";

export default function Favorites(_: RouteFeedProps): JSX.Element {
  return (
    <FeedPage>
      <Feed List={List} query={queryFavorites} name="favorites" />
    </FeedPage>
  );
}
