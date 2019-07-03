import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import IconButton from "@material-ui/core/IconButton";
import FavoriteIcon from "@material-ui/icons/Favorite";
import CircularProgress from "@material-ui/core/CircularProgress";
import { Document } from "../../../../types/document";
import { Bookmark } from "../../../../types/bookmark";
import BookmarkMutation from "../../../apollo/Mutation/Bookmarks/Bookmark";
import { queryFavorites, variables } from "../../../apollo/Query/Feed";

interface Props {
  document: Document;
  onSuccess: (bookmark: Bookmark) => void;
  onError: (message: string) => void;
}

const useStyles = makeStyles(({ palette }) => ({
  button: {
    color: palette.primary.main
  }
}));

export default React.memo(function BookmarkAndFavorite({
  document,
  onSuccess,
  onError
}: Props): JSX.Element {
  const classes = useStyles();
  return (
    <BookmarkMutation
      onCompleted={data => onSuccess(data.bookmarks.bookmark)}
      onError={error => onError(error.message)}
    >
      {(mutate, { loading }) => (
        <IconButton
          aria-label="Bookmark and mark as favorite"
          disabled={loading}
          className={classes.button}
          onClick={() =>
            mutate({
              variables: { url: document.url, isFavorite: true },
              refetchQueries: [
                {
                  query: queryFavorites,
                  variables
                }
              ]
            })
          }
        >
          {!loading && <FavoriteIcon />}
          {loading && <CircularProgress size={16} />}
        </IconButton>
      )}
    </BookmarkMutation>
  );
});
