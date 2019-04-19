import React from "react";
import { withStyles, WithStyles, createStyles } from "@material-ui/core/styles";
import { Query } from "react-apollo";
import Loader from "../../ui/Loader";
import FeedItem from "./Item";
import query from "../../../services/apollo/query/latest.graphql";

import { UserBookmark } from "../../../types/bookmark";

const styles = () =>
  createStyles({
    container: {
      overflow: "hidden",
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

export default withStyles(styles)(function Header({ classes }: Props) {
  return (
    <Query query={query}>
      {({ data, loading, error }) => {
        console.log(data);

        if (loading) {
          return <Loader />;
        }

        return (
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
    </Query>
  );
});
