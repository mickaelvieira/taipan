import React, { useContext } from "react";
import { useMutation } from "@apollo/react-hooks";
import { makeStyles } from "@material-ui/core/styles";
import FavoriteIcon from "@material-ui/icons/Favorite";
import ButtonBase, { ButtonBaseProps } from "../../Button";
import { Document } from "../../../../types/document";
import {
  mutation,
  Data,
  Variables
} from "../../../apollo/Mutation/Bookmarks/Bookmark";
import { FeedsContext, FeedsCacheContext } from "../../../context";
import { SuccessOptions } from ".";

interface Props extends Partial<ButtonBaseProps> {
  document: Document;
  subscriptions?: string[];
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
  subscriptions,
  onSucceed,
  onFail,
  ...rest
}: Props): JSX.Element {
  const classes = useStyles();
  const updater = useContext(FeedsCacheContext);
  const mutator = useContext(FeedsContext);
  const [mutate, { loading }] = useMutation<Data, Variables>(mutation, {
    onCompleted: data => {
      const item = data.bookmarks.add;
      onSucceed({
        updateCache: () => {
          if (updater) {
            updater.bookmark(item);
          }
        },
        undo: () => {
          if (mutator) {
            mutator.unbookmark(item);
          }
        },
        item
      });
    },
    onError: error => onFail(error.message)
  });

  return (
    <ButtonBase
      label="favorites"
      Icon={FavoriteIcon}
      aria-label="Bookmark and mark as favorite"
      disabled={loading}
      className={classes.button}
      onClick={() =>
        mutate({
          variables: { url: document.url, isFavorite: true, subscriptions: subscriptions ? subscriptions : [] }
        })
      }
      {...rest}
    />
  );
});
