import React, { useState } from "react";
import { useMutation } from "@apollo/react-hooks";
import { makeStyles } from "@material-ui/core/styles";
import Chip from "@material-ui/core/Chip";
import DeleteIcon from "@material-ui/icons/Delete";
import PrimaryIcon from "@material-ui/icons/ArrowUpward";
import ListItem from "@material-ui/core/ListItem";
import ListItemText from "@material-ui/core/ListItemText";
import {
  mutation as primaryStatusMutation,
  Variables as PrimaryStatusMutationVariables,
  Data as PrimaryStatusMutationData
} from "../../../apollo/Mutation/User/PrimaryEmail";
import {
  mutation as deleteUserEmailMutation,
  Variables as DeleteUserEmailMutationVariables,
  Data as DeleteUserEmailMutationData
} from "../../../apollo/Mutation/User/DeleteEmail";
import { getErrorMessage } from "../../../apollo/helpers/error";
import { Email } from "../../../../types/users";
import ButtonBase from "../../../ui/Button";
import ConfirmDeleteEmail from "./ConfirmDelete";

const useStyles = makeStyles(() => ({
  button: {
    padding: "12px 6px"
  },
  chip: {
    margin: "0 2px"
  }
}));

interface Props {
  email: Email;
  onDeletionSuccess?: () => void;
  onDeletionFailure: (message: string) => void;
  onStatusSuccess?: () => void;
  onStatusFailure: (message: string) => void;
}

export default function UserEmail({
  email,
  onDeletionFailure,
  onStatusFailure
}: Props): JSX.Element {
  const classes = useStyles();
  const [isShown, setIsShown] = useState(false);
  const [primaryStatus, { loading: isChangingStatus }] = useMutation<
    PrimaryStatusMutationData,
    PrimaryStatusMutationVariables
  >(primaryStatusMutation, {
    onError: error => onStatusFailure(getErrorMessage(error))
  });
  const [deleteEmail, { loading: isDeleting }] = useMutation<
    DeleteUserEmailMutationData,
    DeleteUserEmailMutationVariables
  >(deleteUserEmailMutation, {
    onError: error => onDeletionFailure(getErrorMessage(error))
  });

  return (
    <ListItem disableGutters>
      <ListItemText>{email.value}</ListItemText>
      {!email.isConfirmed && (
        <Chip
          className={classes.chip}
          color="secondary"
          label="Uncomfirmed"
          size="small"
        />
      )}
      {email.isPrimary && (
        <Chip
          className={classes.chip}
          color="primary"
          label="Primary"
          size="small"
        />
      )}
      {!email.isPrimary && (
        <ButtonBase
          label="primary status"
          Icon={PrimaryIcon}
          aria-label="Mark as primary"
          disabled={isChangingStatus}
          onClick={() =>
            primaryStatus({
              variables: { email: email.value }
            })
          }
          iconOnly
          className={classes.button}
        />
      )}
      {!email.isPrimary && (
        <>
          <ButtonBase
            label="remove"
            Icon={DeleteIcon}
            aria-label="Remove email"
            disabled={isDeleting}
            onClick={() => setIsShown(true)}
            iconOnly
            className={classes.button}
          />
          <ConfirmDeleteEmail
            open={isShown}
            onCancel={() => setIsShown(false)}
            onConfirm={() =>
              deleteEmail({
                variables: { email: email.value }
              })
            }
          />
        </>
      )}
    </ListItem>
  );
}
