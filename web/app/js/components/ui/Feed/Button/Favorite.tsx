import React, { useContext } from "react";
import { useMutation } from "@apollo/react-hooks";
import FavoriteIcon from "@material-ui/icons/Favorite";
import ButtonBase, { ButtonBaseProps } from "../../Button";
import { Bookmark } from "../../../../types/bookmark";
import {
  mutation,
  Data,
  Variables,
} from "../../../apollo/Mutation/Bookmarks/Favorite";
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
  const [mutate, { loading }] = useMutation<Data, Variables>(mutation, {
    onCompleted: (data) => {
      const item = data.bookmarks.favorite;
      onSucceed({
        updateCache: () => {
          if (updater) {
            updater.favorite(item);
          }
        },
        undo: () => {
          if (mutator) {
            mutator.unfavorite(item);
          }
        },
        item,
      });
    },
    onError: (error) => onFail(error.message),
  });

  return (
    <ButtonBase
      label="favorites"
      Icon={FavoriteIcon}
      aria-label="Mark as favorite"
      disabled={loading}
      onClick={() =>
        mutate({
          variables: {
            url: bookmark.url,
          },
        })
      }
      {...rest}
    />
  );
});
