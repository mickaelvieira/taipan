import React, { useContext } from "react";
import { useMutation } from "@apollo/react-hooks";
import { makeStyles } from "@material-ui/core/styles";
import FavoriteIcon from "@material-ui/icons/Favorite";
import ButtonBase, { ButtonBaseProps } from "../../Button";
import { Bookmark } from "../../../../types/bookmark";
import {
  mutation,
  Data,
  Variables,
} from "../../../apollo/Mutation/Bookmarks/Unfavorite";
import red from "@material-ui/core/colors/red";
import { FeedsContext, FeedsCacheContext } from "../../../context";
import { SuccessOptions } from ".";

const useStyles = makeStyles({
  button: {
    color: red[800],
  },
});

interface Props extends Partial<ButtonBaseProps> {
  bookmark: Bookmark;
  onSucceed: (options: SuccessOptions) => void;
  onFail: (message: string) => void;
}

export default React.memo(function Unfavorite({
  bookmark,
  onSucceed,
  onFail,
  ...rest
}: Props): JSX.Element {
  const classes = useStyles();
  const updater = useContext(FeedsCacheContext);
  const mutator = useContext(FeedsContext);
  const [mutate, { loading }] = useMutation<Data, Variables>(mutation, {
    onCompleted: (data) => {
      const item = data.bookmarks.unfavorite;
      onSucceed({
        updateCache: () => {
          if (updater) {
            updater.unfavorite(item);
          }
        },
        undo: () => {
          if (mutator) {
            mutator.favorite(item);
          }
        },
        item,
      });
    },
    onError: (error) => onFail(error.message),
  });

  return (
    <ButtonBase
      label="remove"
      Icon={FavoriteIcon}
      aria-label="Remove from favorite"
      className={classes.button}
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
