import React, { useContext } from "react";
import { makeStyles } from "@material-ui/core/styles";
import FavoriteIcon from "@material-ui/icons/Favorite";
import ButtonBase, { ButtonBaseProps } from "../../Button";
import { Document } from "../../../../types/document";
import { Bookmark } from "../../../../types/bookmark";
import BookmarkMutation from "../../../apollo/Mutation/Bookmarks/Bookmark";
import { FeedsContext, FeedsCacheContext } from "../../../context";
import { SuccessOptions } from ".";

interface Props extends Partial<ButtonBaseProps> {
  document: Document;
  onSucceed: (options: SuccessOptions) => void;
  onFail: (message: string) => void;
}

const useStyles = makeStyles(({ palette }) => ({
  button: {
    color: palette.primary.main
  }
}));

export default React.memo(function BookmarkAndFavorite({
  document,
  onSucceed,
  onFail,
  ...rest
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
      onCompleted={data => {
        const item = data.bookmarks.add;
        onSucceed({
          updateCache: getUpdater(item),
          undo: getUndoer(item),
          item
        });
      }}
      onError={error => onFail(error.message)}
    >
      {(mutate, { loading }) => (
        <ButtonBase
          label="favorites"
          Icon={FavoriteIcon}
          aria-label="Bookmark and mark as favorite"
          disabled={loading}
          className={classes.button}
          onClick={() =>
            mutate({
              variables: { url: document.url, isFavorite: true }
            })
          }
          {...rest}
        />
      )}
    </BookmarkMutation>
  );
});
