import React, { PropsWithChildren } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Modal from '@material-ui/core/Modal';
import Paper from "@material-ui/core/Paper";
import { LG_WIDTH} from "../../../constant/feed";

const useStyles = makeStyles(() => ({
  dialog: {},
  paper: {
    position: 'absolute',
    width: LG_WIDTH,
    top: `50%`,
    left: `50%`,
    transform: `translate(-50%, -50%)`,
  }
}));

interface Props {
  isOpen: boolean;
  setIsPanelOpen: (isOpen: boolean) => void;
}

export default function AddForm({
  isOpen,
  setIsPanelOpen,
  children
}: PropsWithChildren<Props>): JSX.Element {
  const classes = useStyles();
  return (
    <Modal open={isOpen} onBackdropClick={() => setIsPanelOpen(false)} >
      <Paper className={classes.paper} elevation={0} square>
        {children}
      </Paper>
    </Modal>
  );
}
