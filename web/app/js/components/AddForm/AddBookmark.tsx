import React, { useState } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import TextField from "@material-ui/core/TextField";
import DialogActions from "@material-ui/core/DialogActions";
import DialogContent from "@material-ui/core/DialogContent";
import DialogTitle from "@material-ui/core/DialogTitle";
import { Bookmark } from "../../types/bookmark";
import CircularProgress from "@material-ui/core/CircularProgress";
import BookmarkMutation, { mutation } from "../apollo/Mutation/Bookmark";
import { query, variables } from "../apollo/Query/ReadingList";

const useStyles = makeStyles({
  title: {
    minWidth: 320
  }
});

interface Props {
  onBookmarkCreated: (bookmark: Bookmark) => void;
  toggleDialog: (status: boolean) => void;
}

export default function AddBookmark({
  onBookmarkCreated,
  toggleDialog
}: Props) {
  const classes = useStyles();
  const [url, setUrl] = useState("");

  return (
    <div>
      <BookmarkMutation
        mutation={mutation}
        onCompleted={({ Bookmark: bookmark }) => {
          setUrl("");
          onBookmarkCreated(bookmark);
        }}
      >
        {(mutate, { loading, error }) => {
          return (
            <>
              <DialogTitle id="form-dialog-title" className={classes.title}>
                Bookmark a document
              </DialogTitle>
              <DialogContent>
                <TextField
                  autoFocus
                  margin="dense"
                  id="bookmark_url"
                  label="URL"
                  placeholder="https://"
                  type="url"
                  value={url}
                  error={!!error}
                  autoComplete="off"
                  autoCapitalize="off"
                  autoCorrect="off"
                  helperText={!error ? "" : error.message}
                  onChange={event => setUrl(event.target.value)}
                  fullWidth
                />
              </DialogContent>
              <DialogActions>
                {loading && <CircularProgress size={16} />}
                <Button
                  onClick={() => toggleDialog(false)}
                  color="primary"
                  disabled={loading}
                >
                  Cancel
                </Button>
                <Button
                  onClick={() =>
                    mutate({
                      variables: { url },
                      refetchQueries: [{ query, variables }]
                    })
                  }
                  color="primary"
                  disabled={loading}
                >
                  Bookmark
                </Button>
              </DialogActions>
            </>
          );
        }}
      </BookmarkMutation>
    </div>
  );
}
