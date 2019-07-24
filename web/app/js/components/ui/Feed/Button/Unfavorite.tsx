import React, { useContext } from "react";
import { makeStyles } from "@material-ui/core/styles";
import IconButton from "@material-ui/core/IconButton";
import FavoriteIcon from "@material-ui/icons/Favorite";
import CircularProgress from "@material-ui/core/CircularProgress";
import { Bookmark } from "../../../../types/bookmark";
import UnfavoriteMutation from "../../../apollo/Mutation/Bookmarks/Unfavorite";
import red from "@material-ui/core/colors/red";
import { Undoer, CacheUpdater } from "../../../../types";
import { FeedsContext, FeedsCacheContext } from "../../../context";

const useStyles = makeStyles({
  button: {
    color: red[800]
  }
});

interface Props {
  bookmark: Bookmark;
  onSuccess: (update: CacheUpdater, undo: Undoer) => void;
  onError: (message: string) => void;
}

export default React.memo(function Unfavorite({
  bookmark,
  onSuccess,
  onError
}: Props): JSX.Element {
  const classes = useStyles();
  const updater = useContext(FeedsCacheContext);
  const mutator = useContext(FeedsContext);
  const getUpdater = (bookmark: Bookmark) => {
    return function() {
      if (updater) {
        updater.unfavorite(bookmark);
      }
    };
  };
  const getUndoer = (bookmark: Bookmark) => {
    return function() {
      if (mutator) {
        mutator.favorite(bookmark);
      }
    };
  };

  return (
    <UnfavoriteMutation
      onCompleted={data =>
        onSuccess(
          getUpdater(data.bookmarks.unfavorite),
          getUndoer(data.bookmarks.unfavorite)
        )
      }
      onError={error => onError(error.message)}
    >
      {(mutate, { loading }) => (
        <IconButton
          aria-label="Remove from favorite"
          className={classes.button}
          disabled={loading}
          onClick={() =>
            mutate({
              variables: {
                url: bookmark.url
              }
            })
          }
        >
          {!loading && <FavoriteIcon />}
          {loading && <CircularProgress size={16} />}
        </IconButton>
      )}
    </UnfavoriteMutation>
  );
});
