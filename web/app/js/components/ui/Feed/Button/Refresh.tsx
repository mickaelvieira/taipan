import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import IconButton from "@material-ui/core/IconButton";
import CachedIcon from "@material-ui/icons/Cached";
import CircularProgress from "@material-ui/core/CircularProgress";
import { Bookmark } from "../../../../types/bookmark";
import BookmarkMutation from "../../../apollo/Mutation/Bookmark";

interface Props {
  bookmark: Bookmark;
  onSuccess: (bookmark: Bookmark) => void;
}

const useStyles = makeStyles(({ palette }) => ({
  button: {
    color: palette.primary.main
  }
}));

export default React.memo(function Bookmark({ bookmark, onSuccess }: Props) {
  const classes = useStyles();
  return (
    <BookmarkMutation onCompleted={data => onSuccess(data.Bookmark)}>
      {(mutate, { loading }) => (
        <IconButton
          aria-label="Refresh"
          disabled={loading}
          className={classes.button}
          onClick={() =>
            mutate({
              variables: { url: bookmark.url }
            })
          }
        >
          {!loading && <CachedIcon />}
          {loading && <CircularProgress size={16} />}
        </IconButton>
      )}
    </BookmarkMutation>
  );
});
