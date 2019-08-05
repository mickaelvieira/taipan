import React, { useState } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Chip from "@material-ui/core/Chip";
import DeleteIcon from "@material-ui/icons/Delete";
import PrimaryIcon from "@material-ui/icons/ArrowUpward";
import ListItem from "@material-ui/core/ListItem";
import ListItemText from "@material-ui/core/ListItemText";
import PrimaryUserEmailMutation from "../../../apollo/Mutation/User/PrimaryEmail";
import DeleteUserEmailMutation from "../../../apollo/Mutation/User/DeleteEmail";
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
  onDeleted: () => void;
}

export default function UserEmail({ email, onDeleted }: Props): JSX.Element {
  const classes = useStyles();
  const [isShown, setIsShown] = useState(false);

  return (
    <ListItem>
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
        <PrimaryUserEmailMutation>
          {(mutate, { loading }) => (
            <ButtonBase
              label="remove"
              Icon={PrimaryIcon}
              aria-label="Mark as primary"
              disabled={loading}
              onClick={() =>
                mutate({
                  variables: { email: email.value }
                })
              }
              iconOnly
              className={classes.button}
            />
          )}
        </PrimaryUserEmailMutation>
      )}
      {!email.isPrimary && (
        <DeleteUserEmailMutation onCompleted={() => onDeleted()}>
          {(mutate, { loading }) => (
            <>
              <ButtonBase
                label="remove"
                Icon={DeleteIcon}
                aria-label="Remove email"
                disabled={loading}
                onClick={() => setIsShown(true)}
                iconOnly
                className={classes.button}
              />
              <ConfirmDeleteEmail
                open={isShown}
                onCancel={() => setIsShown(false)}
                onConfirm={() =>
                  mutate({
                    variables: { email: email.value }
                  })
                }
              />
            </>
          )}
        </DeleteUserEmailMutation>
      )}
    </ListItem>
  );
}
