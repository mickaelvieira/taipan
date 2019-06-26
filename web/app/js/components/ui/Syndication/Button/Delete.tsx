import React, { useState } from "react";
import { makeStyles } from "@material-ui/core/styles";
import IconButton from "@material-ui/core/IconButton";
import DeleteIcon from "@material-ui/icons/DeleteOutline";
import CircularProgress from "@material-ui/core/CircularProgress";
import { Source } from "../../../../types/syndication";
import ConfirmDeletion from "../Confirm/Delete";
import DeleteMutation from "../../../apollo/Mutation/Syndication/Delete";

interface Props {
  source: Source;
  onSuccess: (source: Source) => void;
}

const useStyles = makeStyles(({ palette }) => ({
  button: {
    color: palette.secondary.main
  }
}));

export default React.memo(function Unbookmark({
  source,
  onSuccess
}: Props): JSX.Element {
  const classes = useStyles();
  const [isConfirmVisible, setConfirmVisibility] = useState(false);
  return (
    <DeleteMutation onCompleted={data => onSuccess(data.syndication.delete)}>
      {(mutate, { loading }) => (
        <>
          <IconButton
            aria-label="Remove source"
            disabled={loading}
            onClick={() => setConfirmVisibility(true)}
            className={classes.button}
          >
            {!loading && <DeleteIcon />}
            {loading && <CircularProgress size={16} />}
          </IconButton>
          <ConfirmDeletion
            open={isConfirmVisible}
            onCancel={() => {
              setConfirmVisibility(false);
            }}
            onConfirm={() => {
              setConfirmVisibility(false);
              mutate({
                variables: { url: source.url }
              });
            }}
          />
        </>
      )}
    </DeleteMutation>
  );
});
