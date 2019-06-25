import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import IconButton from "@material-ui/core/IconButton";
import BookmarkIcon from "@material-ui/icons/BookmarkBorderOutlined";
import CircularProgress from "@material-ui/core/CircularProgress";
import { Document } from "../../../../types/document";
import { Bookmark } from "../../../../types/bookmark";
import BookmarkMutation from "../../../apollo/Mutation/Bookmarks/Bookmark";
import { queryReadingList, variables } from "../../../apollo/Query/Feed";

interface Props {
  document: Document;
  onSuccess: (bookmark: Bookmark) => void;
}

const useStyles = makeStyles(({ palette }) => ({
  button: {
    color: palette.primary.main
  }
}));

export default React.memo(function Bookmark({ document, onSuccess }: Props) {
  const classes = useStyles();
  return (
    <BookmarkMutation onCompleted={data => onSuccess(data.bookmarks.bookmark)}>
      {(mutate, { loading }) => (
        <IconButton
          aria-label="Bookmark"
          disabled={loading}
          className={classes.button}
          onClick={() =>
            mutate({
              variables: { url: document.url, isFavorite: false },
              refetchQueries: [
                {
                  query: queryReadingList,
                  variables
                }
              ]
            })
          }
        >
          {!loading && <BookmarkIcon />}
          {loading && <CircularProgress size={16} />}
        </IconButton>
      )}
    </BookmarkMutation>
  );
});
