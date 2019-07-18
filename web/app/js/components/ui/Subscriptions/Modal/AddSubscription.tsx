import React, { useState } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import TextField from "@material-ui/core/TextField";
import DialogActions from "@material-ui/core/DialogActions";
import DialogContent from "@material-ui/core/DialogContent";
import DialogTitle from "@material-ui/core/DialogTitle";
import { Subscription } from "../../../../types/subscription";
import CircularProgress from "@material-ui/core/CircularProgress";
import SubscriptionMutation from "../../../apollo/Mutation/Subscriptions/Subscription";
import Dialog from "@material-ui/core/Dialog";

const useStyles = makeStyles(() => ({
  dialog: {},
  title: {
    minWidth: 320
  }
}));

interface Props {
  isOpen: boolean;
  toggleDialog: (status: boolean) => void;
  onSubscriptionCreated: (subscription: Subscription) => void;
}

export default function AddForm({
  isOpen,
  toggleDialog,
  onSubscriptionCreated
}: Props): JSX.Element {
  const classes = useStyles();
  const [url, setUrl] = useState("");

  return (
    <Dialog
      open={isOpen}
      onClose={() => toggleDialog(false)}
      aria-labelledby="form-dialog-title"
      className={classes.dialog}
    >
      <SubscriptionMutation
        onCompleted={({ subscriptions: { subscription } }) => {
          setUrl("");
          onSubscriptionCreated(subscription);
        }}
      >
        {(mutate, { loading, error }) => {
          return (
            <>
              <DialogTitle id="form-dialog-title" className={classes.title}>
                Add a feed
              </DialogTitle>
              <DialogContent>
                <TextField
                  autoFocus
                  margin="dense"
                  id="feed_url"
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
                      variables: { url }
                    })
                  }
                  color="primary"
                  disabled={loading}
                >
                  Add
                </Button>
              </DialogActions>
            </>
          );
        }}
      </SubscriptionMutation>
    </Dialog>
  );
}