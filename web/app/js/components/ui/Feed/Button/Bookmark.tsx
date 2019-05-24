import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import IconButton from "@material-ui/core/IconButton";
import BookmarkIcon from "@material-ui/icons/BookmarkBorderOutlined";
import CircularProgress from "@material-ui/core/CircularProgress";
import { Document } from "../../../../types/document";
import BookmarkMutation, { mutation } from "../../../apollo/Mutation/Bookmark";

interface Props {
  document: Document;
}

const useStyles = makeStyles(({ palette }) => ({
  button: {
    color: palette.primary.main
  }
}));

export default React.memo(function Bookmark({ document }: Props) {
  const classes = useStyles();
  return (
    <BookmarkMutation mutation={mutation}>
      {(mutate, { loading }) => (
        <IconButton
          aria-label="Bookmark"
          disabled={loading}
          className={classes.button}
          onClick={() =>
            mutate({
              variables: { url: document.url }
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
