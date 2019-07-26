import React, { useContext } from "react";
import FavoriteIcon from "@material-ui/icons/Favorite";
import ButtonBase, { ButtonBaseProps } from "../../Button";
import { Bookmark } from "../../../../types/bookmark";
import FavoriteMutation from "../../../apollo/Mutation/Bookmarks/Favorite";
import { FeedsContext, FeedsCacheContext } from "../../../context";
import { SuccessOptions } from ".";

interface Props extends Partial<ButtonBaseProps> {
  bookmark: Bookmark;
  onSucceed: (options: SuccessOptions) => void;
  onFail: (message: string) => void;
}

export default React.memo(function Favorite({
  bookmark,
  onSucceed,
  onFail,
  ...rest
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
      onCompleted={data => {
        const item = data.bookmarks.favorite;
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
          aria-label="Mark as favorite"
          disabled={loading}
          onClick={() =>
            mutate({
              variables: {
                url: bookmark.url
              }
            })
          }
          {...rest}
        />
      )}
    </FavoriteMutation>
  );
});
