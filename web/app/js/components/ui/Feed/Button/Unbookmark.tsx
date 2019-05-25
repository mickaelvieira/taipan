import React, { useState } from "react";
import { makeStyles } from "@material-ui/core/styles";
import IconButton from "@material-ui/core/IconButton";
import DeleteIcon from "@material-ui/icons/DeleteRounded";
import CircularProgress from "@material-ui/core/CircularProgress";
import { Bookmark } from "../../../../types/bookmark";
import ConfirmUnbookmark from "../Confirm/Unbookmark";
import UnbookmarkMutation, {
  mutation
} from "../../../apollo/Mutation/Unbookmark";

interface Props {
  bookmark: Bookmark;
}

const useStyles = makeStyles(({ palette }) => ({
  button: {
    color: palette.primary.main
  }
}));

export default React.memo(function Unbookmark({ bookmark }: Props) {
  const classes = useStyles();
  const [isConfirmVisible, setConfirmVisibility] = useState(false);
  return (
    <UnbookmarkMutation mutation={mutation}>
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
                variables: { url: bookmark.url }
              });
            }}
          />
        </>
      )}
    </UnbookmarkMutation>
  );
});
