import React, { useState } from "react";
import { makeStyles } from "@material-ui/core/styles";
import IconButton from "@material-ui/core/IconButton";
import DeleteIcon from "@material-ui/icons/DeleteOutline";
import CircularProgress from "@material-ui/core/CircularProgress";
import { Bookmark } from "../../../../types/bookmark";
import { Document } from "../../../../types/document";
import ConfirmUnbookmark from "../Confirm/Unbookmark";
import UnbookmarkMutation from "../../../apollo/Mutation/Bookmarks/Unbookmark";
import {
  queryReadingList,
  queryFavorites,
  variables
} from "../../../apollo/Query/Feed";

interface Props {
  bookmark: Bookmark;
  onSuccess: (document: Document) => void;
}

const useStyles = makeStyles(({ palette }) => ({
  button: {
    color: palette.primary.main
  }
}));

export default React.memo(function Unbookmark({
  bookmark,
  onSuccess
}: Props): JSX.Element {
  const classes = useStyles();
  const [isConfirmVisible, setConfirmVisibility] = useState(false);
  return (
    <UnbookmarkMutation
      onCompleted={data => onSuccess(data.bookmarks.unbookmark)}
    >
      {(mutate, { loading }) => (
        <>
          <IconButton
            aria-label="Remove bookmark"
            disabled={loading}
            onClick={() => setConfirmVisibility(true)}
            className={classes.button}
          >
            {!loading && <DeleteIcon />}
            {loading && <CircularProgress size={16} />}
          </IconButton>
          <ConfirmUnbookmark
            open={isConfirmVisible}
            onCancel={() => {
              setConfirmVisibility(false);
            }}
            onConfirm={() => {
              setConfirmVisibility(false);
              mutate({
                variables: { url: bookmark.url },
                refetchQueries: [
                  {
                    query: queryFavorites,
                    variables
                  },
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
    </UnbookmarkMutation>
  );
});
