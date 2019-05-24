import React from "react";
import Button from "@material-ui/core/Button";
import DialogTitle from "@material-ui/core/DialogTitle";
import DialogActions from "@material-ui/core/DialogActions";
import Dialog from "@material-ui/core/Dialog";

export interface Props {
  open: boolean;
  onConfirm: () => void;
  onCancel: () => void;
}

export default function ConfirmUnbookmark({
  onConfirm,
  onCancel,
  open
}: Props) {
  return (
    <Dialog
      disableBackdropClick
      disableEscapeKeyDown
      maxWidth="xs"
      aria-labelledby="confirmation-dialog-title"
      open={open}
    >
      <DialogTitle>Do you want to remove this bookmark?</DialogTitle>
      <DialogActions>
        <Button onClick={onCancel} color="primary">
          Cancel
        </Button>
        <Button onClick={onConfirm} color="primary">
          Remove
        </Button>
      </DialogActions>
    </Dialog>
  );
}
