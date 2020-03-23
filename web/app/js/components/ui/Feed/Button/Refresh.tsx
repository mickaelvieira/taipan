import React from "react";
import { useMutation } from "@apollo/react-hooks";
import { makeStyles } from "@material-ui/core/styles";
import CachedIcon from "@material-ui/icons/Cached";
import ButtonBase, { ButtonBaseProps } from "../../Button";
import { Bookmark } from "../../../../types/bookmark";
import {
  mutation,
  variables,
  Data,
  Variables,
} from "../../../apollo/Mutation/Bookmarks/Create";

interface Props extends Partial<ButtonBaseProps> {
  bookmark: Bookmark;
  onSucceed: (bookmark: Bookmark) => void;
  onFail: (message: string) => void;
}

const useStyles = makeStyles(({ palette }) => ({
  button: {
    color: palette.primary.main,
  },
}));

export default React.memo(function Refresh({
  bookmark,
  onSucceed,
  onFail,
  ...rest
}: Props): JSX.Element {
  const classes = useStyles();
  const [mutate, { loading }] = useMutation<Data, Variables>(mutation, {
    onCompleted: (data) => onSucceed(data.bookmarks.create),
    onError: (error) => onFail(error.message),
  });
  return (
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
            url: bookmark.url,
          },
        })
      }
      {...rest}
    />
  );
});
