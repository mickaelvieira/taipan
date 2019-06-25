import React, { useState } from "react";
import { makeStyles } from "@material-ui/core/styles";
import IconButton from "@material-ui/core/IconButton";
import FavoriteIcon from "@material-ui/icons/Favorite";
import CircularProgress from "@material-ui/core/CircularProgress";
import { Bookmark } from "../../../../types/bookmark";
import UnfavoriteMutation from "../../../apollo/Mutation/Bookmarks/Unfavorite";
import { queryReadingList, variables } from "../../../apollo/Query/Feed";
import ConfirmUnfavorite from "../Confirm/Unfavorite";
import red from "@material-ui/core/colors/red";

const useStyles = makeStyles({
  button: {
    color: red[800]
  }
});

interface Props {
  bookmark: Bookmark;
  onSuccess: (bookmark: Bookmark) => void;
}

export default React.memo(function Unfavorite({ bookmark, onSuccess }: Props) {
  const classes = useStyles();
  const [isConfirmVisible, setConfirmVisibility] = useState(false);

  return (
    <UnfavoriteMutation
      onCompleted={data => onSuccess(data.bookmarks.unfavorite)}
    >
      {(mutate, { loading }) => (
        <>
          <IconButton
            aria-label="Remove from favorite"
            className={classes.button}
            disabled={loading}
            onClick={() => setConfirmVisibility(true)}
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
                variables: {
                  url: bookmark.url
                },
                refetchQueries: [
                  {
                    query: queryReadingList,
                    variables
                  }
                ]
              });
            }}
          />
        </>
      )}
    </UnfavoriteMutation>
  );
});
