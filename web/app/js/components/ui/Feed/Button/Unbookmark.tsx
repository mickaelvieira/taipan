import React, { useContext } from "react";
import { makeStyles } from "@material-ui/core/styles";
import IconButton from "@material-ui/core/IconButton";
import DeleteIcon from "@material-ui/icons/DeleteOutline";
import CircularProgress from "@material-ui/core/CircularProgress";
import { Bookmark } from "../../../../types/bookmark";
import { Document } from "../../../../types/document";
import UnbookmarkMutation from "../../../apollo/Mutation/Bookmarks/Unbookmark";
import { FeedsContext, FeedsCacheContext } from "../../../context";
import { Undoer, CacheUpdater } from "../../../../types";

interface Props {
  bookmark: Bookmark;
  onSuccess: (update: CacheUpdater, undo: Undoer) => void;
  onError: (message: string) => void;
}

const useStyles = makeStyles(({ palette }) => ({
  button: {
    color: palette.primary.main
  }
}));

export default React.memo(function Unbookmark({
  bookmark,
  onSuccess,
  onError
}: Props): JSX.Element {
  const classes = useStyles();
  const updater = useContext(FeedsCacheContext);
  const mutator = useContext(FeedsContext);
  const getUpdater = (document: Document) => {
    return function() {
      if (updater) {
        updater.unbookmark(document);
      }
    };
  };
  const getUndoer = (document: Document) => {
    return function() {
      if (mutator) {
        mutator.bookmark(document, bookmark.isFavorite);
      }
    };
  };

  return (
    <UnbookmarkMutation
      onCompleted={data =>
        onSuccess(
          getUpdater(data.bookmarks.remove),
          getUndoer(data.bookmarks.remove)
        )
      }
      onError={error => onError(error.message)}
    >
      {(mutate, { loading }) => (
        <IconButton
          aria-label="Remove bookmark"
          disabled={loading}
          onClick={() =>
            mutate({
              variables: { url: bookmark.url }
            })
          }
          className={classes.button}
        >
          {!loading && <DeleteIcon />}
          {loading && <CircularProgress size={16} />}
        </IconButton>
      )}
    </UnbookmarkMutation>
  );
});
