import React, { useState } from "react";
import { makeStyles } from "@material-ui/core/styles";
import IconButton from "@material-ui/core/IconButton";
import DeleteIcon from "@material-ui/icons/DeleteOutline";
import CircularProgress from "@material-ui/core/CircularProgress";
import { Source } from "../../../../types/syndication";
import ConfirmDeletion from "../Confirm/Delete";
import DeleteSourceMutation from "../../../apollo/Mutation/Syndication/Delete";

interface Props {
  source: Source;
  onSuccess: (source: Source) => void;
  onError: (message: string) => void;
}

const useStyles = makeStyles(({ palette }) => ({
  button: {
    color: palette.secondary.main
  }
}));

export default React.memo(function DeleteButton({
  source,
  onSuccess,
  onError
}: Props): JSX.Element {
  const classes = useStyles();
  const [isConfirmVisible, setConfirmVisibility] = useState(false);
  return (
    <DeleteSourceMutation
      onCompleted={data => onSuccess(data.syndication.delete)}
      onError={error => onError(error.message)}
    >
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
    </DeleteSourceMutation>
  );
});
