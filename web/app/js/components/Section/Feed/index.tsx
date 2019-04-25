import React from "react";
import { withStyles, WithStyles, createStyles } from "@material-ui/core/styles";
import Loader from "../../ui/Loader";
import FeedItem from "./Item";
import LatestBookmarksQuery from "../../apollo/Query/LatestBookmarks";
import query from "../../../services/apollo/query/latest.graphql";
import { UserBookmark } from "../../../types/bookmark";

const styles = () =>
  createStyles({
    container: {
      overflow: "auto",
      display: "flex",
      height: 600,
      width: 424
    },
    list: {
      listStyleType: "none",
      display: "flex",
      flexDirection: "row",
      padding: 0
    },
    listItem: {
      display: "flex",
      flex: 1,
      flexDirection: "column",
      minWidth: 424,
      maxWidth: 424
    }
  });

interface Props extends WithStyles<typeof styles> {
  bookmark: UserBookmark;
}

export default withStyles(styles)(function Feed({ classes }: Props) {
  return (
    <LatestBookmarksQuery query={query} variables={{ limit: 30 }}>
      {({ data, loading, error }) => {
        if (loading) {
          return <Loader />;
        }

        return !data ? null : (
          <div className={classes.container}>
            <ul className={classes.list}>
              {data.GetLatestBookmarks.results.map((bookmark: UserBookmark) => {
                return (
                  <li className={classes.listItem} key={bookmark.id}>
                    <FeedItem bookmark={bookmark} />
                  </li>
                );
              })}
            </ul>
          </div>
        );
      }}
    </LatestBookmarksQuery>
  );
});
