import React from "react";
import { withStyles, WithStyles, createStyles } from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import TextField from "@material-ui/core/TextField";
import Dialog from "@material-ui/core/Dialog";
import DialogActions from "@material-ui/core/DialogActions";
import DialogContent from "@material-ui/core/DialogContent";
import DialogTitle from "@material-ui/core/DialogTitle";
import { UserBookmark } from "../../types/bookmark";

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
  classes
}: Props) {
  return (
    <Dialog
      open={isOpen}
      onClose={() => toggleDialog(false)}
      aria-labelledby="form-dialog-title"
      className={classes.dialog}
    >
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
          fullWidth
        />
      </DialogContent>
      <DialogActions>
        <Button onClick={() => toggleDialog(false)} color="primary">
          Cancel
        </Button>
        <Button onClick={() => toggleDialog(false)} color="primary">
          Bookmark
        </Button>
      </DialogActions>
    </Dialog>
  );
});
