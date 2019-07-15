import React, { useState } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Dialog from "@material-ui/core/Dialog";
import Button from "@material-ui/core/Button";
import TextField from "@material-ui/core/TextField";
import DialogActions from "@material-ui/core/DialogActions";
import DialogContent from "@material-ui/core/DialogContent";
import DialogTitle from "@material-ui/core/DialogTitle";
import FormGroup from "@material-ui/core/FormGroup";
import FormControlLabel from "@material-ui/core/FormControlLabel";
import Checkbox from "@material-ui/core/Checkbox";
import { Bookmark } from "../../../../types/bookmark";
import CircularProgress from "@material-ui/core/CircularProgress";
import CreateBookmarkMutation, { variables } from "../../../apollo/Mutation/Bookmarks/Create";

const useStyles = makeStyles(() => ({
  dialog: {},
  title: {
    minWidth: 320
  }
}));

interface Props {
  onBookmarkCreated: (bookmark: Bookmark) => void;
  toggleDialog: (status: boolean) => void;
}

interface Props {
  isOpen: boolean;
  toggleDialog: (status: boolean) => void;
  onBookmarkCreated: (bookmark: Bookmark) => void;
}

// @TODO Adds the newly created bookmark to reading list or favorite
// Add a checkbox to add it directly to favorites
export default function AddForm({
  isOpen,
  toggleDialog,
  onBookmarkCreated
}: Props): JSX.Element {
  const classes = useStyles();
  const [url, setUrl] = useState("");
  const [withFeeds, setWithFeeds] = useState(variables.withFeeds);
  const [isFavorite, setIsFavorite] = useState(variables.isFavorite);

  return (
    <Dialog
      open={isOpen}
      onClose={() => toggleDialog(false)}
      aria-labelledby="form-dialog-title"
      className={classes.dialog}
    >
      <CreateBookmarkMutation
        onCompleted={({ bookmarks: { create: bookmark } }) => {
          setUrl("");
          onBookmarkCreated(bookmark);
        }}
      >
        {(mutate, { loading, error }) => {
          return (
            <form>
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
                <FormGroup>
                  <FormControlLabel
                    control={
                      <Checkbox
                        checked={isFavorite}
                        onChange={() => setIsFavorite(!isFavorite)}
                        value="1"
                        color="primary"
                      />
                    }
                    label="Add to favorites"
                  />
                </FormGroup>
                <FormGroup>
                  <FormControlLabel
                    control={
                      <Checkbox
                        checked={withFeeds}
                        onChange={() => setWithFeeds(!withFeeds)}
                        value="1"
                        color="primary"
                      />
                    }
                    label="Parse RSS feeds"
                  />
                </FormGroup>
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
                      variables: { url, withFeeds, isFavorite }
                    })
                  }
                  color="primary"
                  disabled={loading}
                >
                  Bookmark
                </Button>
              </DialogActions>
            </form>
          );
        }}
      </CreateBookmarkMutation>
    </Dialog>
  );
}
