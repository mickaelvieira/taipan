import React, { useContext } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Typography from "@material-ui/core/Typography";
import Button from "@material-ui/core/Button";
import { Document } from "../../../types/document";
import Item from "./Item";
import { ListProps } from "../../ui/Feed/Feed";
import { UserContext } from "../../context";
import Latest from "./Latest";
import FeedItem from "../../ui/Feed/Item/Item";
import EmptyFeed from "../../ui/Feed/Empty";
import { RouterLink } from "../../ui/Link";
import Emoji from "../../ui/Emoji";

const useStyles = makeStyles(({ spacing }) => ({
  message: {
    margin: spacing(2)
  },
  subscriptions: {
    display: "flex",
    alignItems: "center",
    margin: spacing(2)
  },
  links: {
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
    margin: spacing(2)
  }
}));

export default React.memo(function DocumentList({
  results,
  firstId,
  lastId
}: ListProps): JSX.Element {
  const user = useContext(UserContext);
  const classes = useStyles();

  return (
    <>
      <Latest firstId={firstId} lastId={lastId} />
      {results.length === 0 && (
        <>
          <EmptyFeed>
            <Typography className={classes.message}>
              There isn&apos;t any news today.
            </Typography>
            {user && user.stats && user.stats.subscriptions === 0 && (
              <>
                <Typography className={classes.subscriptions}>
                  But, you haven&apos;t subscribed to any feeds.
                </Typography>
                <div className={classes.links}>
                  <Button
                    to="/syndication"
                    variant="contained"
                    component={RouterLink}
                    color="primary"
                  >
                    Discover <Emoji emoji=":heart_eyes:" />
                  </Button>
                </div>
              </>
            )}
          </EmptyFeed>
        </>
      )}
      {results.map(result => (
        <FeedItem key={result.id}>
          <Item document={result as Document} />
        </FeedItem>
      ))}
    </>
  );
});
