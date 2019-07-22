import React, { useState } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Link from "@material-ui/core/Link";
import Checkbox from "@material-ui/core/Checkbox";
import Typography from "@material-ui/core/Typography";
import List from "@material-ui/core/List";
import ListItem from "@material-ui/core/ListItem";
import CardMedia from "@material-ui/core/CardMedia";
import Button from "@material-ui/core/Button";
import { Document } from "../../../../types/document";
import { Bookmark } from "../../../../types/bookmark";
import Title from "../Item/Title";
import Description from "../Item/Description";
import { getImageWidth } from "../../../../helpers/image";
import BookmarkMutation from "../../../apollo/Mutation/Bookmarks/Bookmark";

const useStyles = makeStyles(({ breakpoints }) => ({
  form: {
    display: "flex",
    flexDirection: "column"
  },
  media: {
    marginTop: 16,
    marginBottom: 16,
    backgroundSize: "cover",
    minHeight: `calc(${getImageWidth("sm")}px * 9 / 16)`,
    [breakpoints.up("md")]: {
      minHeight: `calc(${getImageWidth("sm")}px * 9 / 16)`
    }
  },
  item: {
    padding: 0
  },
  link: {
    paddingTop: 10,
    paddingLeft: 9,
    paddingBottom: 9
  },
  actions: {
    display: "flex",
    flexDirection: "row",
    justifyContent: "flex-end",
    alignItems: "center"
  },
  addto: {
    padding: "8px 6px"
  }
}));

interface Props {
  document: Document;
  onBookmarkCreated: (bookmark: Bookmark) => void;
  onRemoveBookmark: () => void;
}

export default function FormDocument({
  document,
  onRemoveBookmark,
  onBookmarkCreated
}: Props): JSX.Element {
  const classes = useStyles();
  const [subscriptions, setSubscriptions] = useState<string[]>([]);
  const syndication = document.syndication ? document.syndication : [];
  const { url } = document;

  return (
    <BookmarkMutation
      onCompleted={({ bookmarks: { add: bookmark } }) =>
        onBookmarkCreated(bookmark)
      }
    >
      {(mutate, { loading }) => {
        return (
          <form className={classes.form}>
            <Title item={document} />
            <Description item={document} />
            {document.image && (
              <CardMedia
                className={classes.media}
                image={document.image.url}
                title={document.title}
              />
            )}
            {syndication.length > 0 && (
              <>
                <Typography>
                  We also found the following RSS feeds. Do you want to
                  subscribe to them?
                </Typography>
                <List>
                  {syndication.map(source => (
                    <ListItem
                      className={classes.item}
                      alignItems="flex-start"
                      key={source.id}
                    >
                      <Checkbox
                        checked={subscriptions.includes(source.url)}
                        onClick={() => {
                          let sub = subscriptions.includes(source.url)
                            ? subscriptions.filter(url => url != source.url)
                            : [source.url, ...subscriptions];
                          setSubscriptions(sub);
                        }}
                      />
                      <Link
                        underline="none"
                        href={source.url}
                        title={source.title}
                        target="_blank"
                        rel="noopener"
                        className={classes.link}
                      >
                        {source.title && source.title != "wordpress feed"
                          ? source.title
                          : source.url}
                      </Link>
                    </ListItem>
                  ))}
                </List>
              </>
            )}
            <div className={classes.actions}>
              <Typography variant="button" className={classes.addto}>
                Add to
              </Typography>
              <Button
                onClick={() =>
                  mutate({
                    variables: { url, subscriptions, isFavorite: false }
                  })
                }
                color="primary"
                disabled={loading}
              >
                Reading list
              </Button>
              <Button
                onClick={() =>
                  mutate({
                    variables: { url, subscriptions, isFavorite: true }
                  })
                }
                color="primary"
                disabled={loading}
              >
                Favorites
              </Button>
              <Typography variant="button" className={classes.addto}>
                or
              </Typography>
              <Button
                onClick={() => onRemoveBookmark()}
                color="secondary"
                disabled={loading}
              >
                Cancel
              </Button>
            </div>
          </form>
        );
      }}
    </BookmarkMutation>
  );
}
