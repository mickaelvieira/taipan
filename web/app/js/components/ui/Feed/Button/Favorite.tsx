import React from "react";
import { withStyles, WithStyles, createStyles } from "@material-ui/core/styles";
import IconButton from "@material-ui/core/IconButton";
import FavoriteIcon from "@material-ui/icons/Favorite";
import CircularProgress from "@material-ui/core/CircularProgress";
import { Bookmark } from "../../../../types/bookmark";
import FavoriteMutation, { mutation } from "../../../apollo/Mutation/Favorite";
import red from "@material-ui/core/colors/red";

const styles = () =>
  createStyles({
    inactive: {},
    active: {
      color: red[800]
    }
  });

interface Props extends WithStyles<typeof styles> {
  bookmark: Bookmark;
}

export default withStyles(styles)(
  React.memo(function Favorite({ bookmark, classes }: Props) {
    return (
      <FavoriteMutation mutation={mutation}>
        {(mutate, { loading }) => (
          <IconButton
            aria-label={
              bookmark.isRead ? "Mark as favorite" : "Remove from favorite"
            }
            className={bookmark.isRead ? classes.active : classes.inactive}
            disabled={loading}
            onClick={() =>
              mutate({
                variables: { url: bookmark.url, isFavorite: !bookmark.isRead }
              })
            }
          >
            {!loading && <FavoriteIcon />}
            {loading && <CircularProgress size={16} />}
          </IconButton>
        )}
      </FavoriteMutation>
    );
  })
);
