import React, { useState } from "react";
import { withStyles, WithStyles, createStyles } from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import TextField from "@material-ui/core/TextField";
import Dialog from "@material-ui/core/Dialog";
import DialogActions from "@material-ui/core/DialogActions";
import DialogContent from "@material-ui/core/DialogContent";
import DialogTitle from "@material-ui/core/DialogTitle";
import { UserBookmark } from "../../types/bookmark";
import CircularProgress from "@material-ui/core/CircularProgress";
import CreateBookmarkMutation, {
  mutation
} from "../apollo/Mutation/CreateBookmark";
import { query, variables } from "../apollo/Query/LatestBookmarks";

const styles = () =>
  createStyles({
    dialog: {},
    title: {
      minWidth: 320
    }
  });

interface Props extends WithStyles<typeof styles> {
  isOpen: boolean;
  toggleDialog: (status: boolean) => void;
  onBookmarkCreated: (bookmark: UserBookmark) => void;
}

export default withStyles(styles)(function AddBookmark({
  isOpen,
  toggleDialog,
  onBookmarkCreated,
  classes
}: Props) {
  const [url, setUrl] = useState("");

  return (
    <Dialog
      open={isOpen}
      onClose={() => toggleDialog(false)}
      aria-labelledby="form-dialog-title"
      className={classes.dialog}
    >
      <CreateBookmarkMutation
        mutation={mutation}
        onCompleted={({ CreateBookmark: bookmark }) => {
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
      </CreateBookmarkMutation>
    </Dialog>
  );
});
