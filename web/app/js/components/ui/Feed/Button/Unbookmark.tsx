import React, { useContext } from "react";
import { useMutation } from "@apollo/react-hooks";
import { makeStyles } from "@material-ui/core/styles";
import DeleteIcon from "@material-ui/icons/Delete";
import ButtonBase, { ButtonBaseProps } from "../../Button";
import { Bookmark } from "../../../../types/bookmark";
import { Document } from "../../../../types/document";
import {
  mutation,
  Data,
  Variables
} from "../../../apollo/Mutation/Bookmarks/Unbookmark";
import { FeedsContext, FeedsCacheContext } from "../../../context";
import { SuccessOptions } from ".";

interface Props extends Partial<ButtonBaseProps> {
  bookmark: Bookmark;
  onSucceed: (options: SuccessOptions) => void;
  onFail: (message: string) => void;
}

const useStyles = makeStyles(({ palette }) => ({
  button: {
    color: palette.primary.main
  }
}));

export default React.memo(function Unbookmark({
  bookmark,
  onSucceed,
  onFail,
  ...rest
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
  const [mutate, { loading }] = useMutation<Data, Variables>(mutation, {
    onCompleted: data => {
      const item = data.bookmarks.remove;
      onSucceed({
        updateCache: getUpdater(item),
        undo: getUndoer(item),
        item
      });
    },
    onError: error => onFail(error.message)
  });

  return (
    <ButtonBase
      label="remove"
      Icon={DeleteIcon}
      aria-label="Remove bookmark"
      disabled={loading}
      onClick={() =>
        mutate({
          variables: { url: bookmark.url }
        })
      }
      className={classes.button}
      {...rest}
    />
  );
});
