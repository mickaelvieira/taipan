import React from "react";
import { withStyles, WithStyles, createStyles } from "@material-ui/core/styles";
import IconButton from "@material-ui/core/IconButton";
import BookmarkIcon from "@material-ui/icons/BookmarkBorderOutlined";
import CircularProgress from "@material-ui/core/CircularProgress";
import { Document } from "../../../../types/document";
import BookmarkMutation, { mutation } from "../../../apollo/Mutation/Bookmark";

const styles = () => createStyles({});

interface Props extends WithStyles<typeof styles> {
  document: Document;
}

export default withStyles(styles)(
  React.memo(function Bookmark({ document }: Props) {
    return (
      <BookmarkMutation mutation={mutation}>
        {(mutate, { loading }) => (
          <IconButton
            aria-label="Bookmark"
            disabled={loading}
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
  })
);
