import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import CachedIcon from "@material-ui/icons/Cached";
import ButtonBase, { ButtonBaseProps } from "../../Button";
import { Bookmark } from "../../../../types/bookmark";
import CreateBookmarkMutation, {
  variables
} from "../../../apollo/Mutation/Bookmarks/Create";

interface Props extends Partial<ButtonBaseProps> {
  bookmark: Bookmark;
  onSucceed: (bookmark: Bookmark) => void;
  onFail: (message: string) => void;
}

const useStyles = makeStyles(({ palette }) => ({
  button: {
    color: palette.primary.main
  }
}));

export default React.memo(function Refresh({
  bookmark,
  onSucceed,
  onFail,
  ...rest
}: Props): JSX.Element {
  const classes = useStyles();
  return (
    <CreateBookmarkMutation
      onCompleted={data => onSucceed(data.bookmarks.create)}
      onError={error => onFail(error.message)}
    >
      {(mutate, { loading }) => (
        <ButtonBase
          label="refresh"
          Icon={CachedIcon}
          isLoading={loading}
          aria-label="Refresh"
          disabled={loading}
          className={classes.button}
          onClick={() =>
            mutate({
              variables: {
                ...variables,
                isFavorite: bookmark.isFavorite,
                url: bookmark.url
              }
            })
          }
          {...rest}
        />
      )}
    </CreateBookmarkMutation>
  );
});
