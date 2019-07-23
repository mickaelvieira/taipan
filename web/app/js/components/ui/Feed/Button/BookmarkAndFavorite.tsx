import React, { useContext } from "react";
import { makeStyles } from "@material-ui/core/styles";
import IconButton from "@material-ui/core/IconButton";
import FavoriteIcon from "@material-ui/icons/Favorite";
import CircularProgress from "@material-ui/core/CircularProgress";
import { Document } from "../../../../types/document";
import { Bookmark } from "../../../../types/bookmark";
import BookmarkMutation from "../../../apollo/Mutation/Bookmarks/Bookmark";
import { FeedsContext, FeedsCacheContext } from "../../../context";
import { Undoer, CacheUpdater } from "../../../../types";

interface Props {
  document: Document;
  onSuccess: (update: CacheUpdater, undo: Undoer) => void;
  onError: (message: string) => void;
}

const useStyles = makeStyles(({ palette }) => ({
  button: {
    color: palette.primary.main
  }
}));

export default React.memo(function BookmarkAndFavorite({
  document,
  onSuccess,
  onError
}: Props): JSX.Element {
  const classes = useStyles();
  const updater = useContext(FeedsCacheContext);
  const mutator = useContext(FeedsContext);
  const getUpdater = (bookmark: Bookmark) => {
    return function() {
      if (updater) {
        updater.bookmark(bookmark);
      }
    };
  };
  const getUndoer = (bookmark: Bookmark) => {
    return function() {
      if (mutator) {
        mutator.unbookmark(bookmark);
      }
    };
  };

  return (
    <BookmarkMutation
      onCompleted={data =>
        onSuccess(getUpdater(data.bookmarks.add), getUndoer(data.bookmarks.add))
      }
      onError={error => onError(error.message)}
    >
      {(mutate, { loading }) => (
        <IconButton
          aria-label="Bookmark and mark as favorite"
          disabled={loading}
          className={classes.button}
          onClick={() =>
            mutate({
              variables: { url: document.url, isFavorite: true }
            })
          }
        >
          {!loading && <FavoriteIcon />}
          {loading && <CircularProgress size={16} />}
        </IconButton>
      )}
    </BookmarkMutation>
  );
});
