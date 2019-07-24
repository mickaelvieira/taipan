import React, { useContext } from "react";
import IconButton from "@material-ui/core/IconButton";
import FavoriteIcon from "@material-ui/icons/Favorite";
import CircularProgress from "@material-ui/core/CircularProgress";
import { Bookmark } from "../../../../types/bookmark";
import FavoriteMutation from "../../../apollo/Mutation/Bookmarks/Favorite";
import { Undoer, CacheUpdater } from "../../../../types";
import { FeedsContext, FeedsCacheContext } from "../../../context";

interface Props {
  bookmark: Bookmark;
  onSuccess: (update: CacheUpdater, undo: Undoer) => void;
  onError: (message: string) => void;
}

export default React.memo(function Favorite({
  bookmark,
  onSuccess,
  onError
}: Props): JSX.Element {
  const updater = useContext(FeedsCacheContext);
  const mutator = useContext(FeedsContext);
  const getUpdater = (bookmark: Bookmark) => {
    return function() {
      if (updater) {
        updater.favorite(bookmark);
      }
    };
  };
  const getUndoer = (bookmark: Bookmark) => {
    return function() {
      if (mutator) {
        mutator.unfavorite(bookmark);
      }
    };
  };

  return (
    <FavoriteMutation
      onCompleted={data =>
        onSuccess(
          getUpdater(data.bookmarks.favorite),
          getUndoer(data.bookmarks.favorite)
        )
      }
      onError={error => onError(error.message)}
    >
      {(mutate, { loading }) => (
        <>
          <IconButton
            aria-label="Mark as favorite"
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
        </>
      )}
    </FavoriteMutation>
  );
});
