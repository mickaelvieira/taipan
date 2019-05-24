import React from "react";
import { withStyles, WithStyles, createStyles } from "@material-ui/core/styles";
import IconButton from "@material-ui/core/IconButton";
import CachedIcon from "@material-ui/icons/Cached";
import CircularProgress from "@material-ui/core/CircularProgress";
import { Bookmark } from "../../../../types/bookmark";
import BookmarkMutation, { mutation } from "../../../apollo/Mutation/Bookmark";

const styles = () => createStyles({});

interface Props extends WithStyles<typeof styles> {
  bookmark: Bookmark;
}

export default withStyles(styles)(
  React.memo(function Bookmark({ bookmark }: Props) {
    return (
      <BookmarkMutation mutation={mutation}>
        {(mutate, { loading }) => (
          <IconButton
            aria-label="Refresh"
            disabled={loading}
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
  })
);
