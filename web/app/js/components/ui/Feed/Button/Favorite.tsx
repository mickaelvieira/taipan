import React, { useState } from "react";
import { makeStyles } from "@material-ui/core/styles";
import IconButton from "@material-ui/core/IconButton";
import FavoriteIcon from "@material-ui/icons/Favorite";
import CircularProgress from "@material-ui/core/CircularProgress";
import { Bookmark } from "../../../../types/bookmark";
import FavoriteMutation from "../../../apollo/Mutation/Favorite";
import ConfirmUnfavorite from "../Confirm/Unfavorite";
import red from "@material-ui/core/colors/red";

const useStyles = makeStyles({
  inactive: {},
  active: {
    color: red[800]
  }
});

interface Props {
  bookmark: Bookmark;
  onSuccess: (bookmark: Bookmark) => void;
}

export default React.memo(function Favorite({ bookmark, onSuccess }: Props) {
  const classes = useStyles();
  const [isConfirmVisible, setConfirmVisibility] = useState(false);

  return (
    <FavoriteMutation
      onCompleted={data => onSuccess(data.ChangeBookmarkReadStatus)}
    >
      {(mutate, { loading }) => (
        <>
          <IconButton
            aria-label={
              bookmark.isRead ? "Remove from favorite" : "Mark as favorite"
            }
            className={bookmark.isRead ? classes.active : classes.inactive}
            disabled={loading}
            onClick={() => {
              if (bookmark.isRead) {
                setConfirmVisibility(true);
              } else {
                mutate({
                  variables: {
                    url: bookmark.url,
                    isFavorite: !bookmark.isRead
                  }
                });
              }
            }}
          >
            {!loading && <FavoriteIcon />}
            {loading && <CircularProgress size={16} />}
          </IconButton>
          <ConfirmUnfavorite
            open={isConfirmVisible}
            onCancel={() => {
              setConfirmVisibility(false);
            }}
            onConfirm={() => {
              setConfirmVisibility(false);
              mutate({
                variables: { url: bookmark.url, isFavorite: !bookmark.isRead }
              });
            }}
          />
        </>
      )}
    </FavoriteMutation>
  );
});
