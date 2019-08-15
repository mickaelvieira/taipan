import React, { useState, useContext } from "react";
import { useMutation } from "@apollo/react-hooks";
import { makeStyles } from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import Card from "../Card";
import CardContent from "@material-ui/core/CardContent";
import CardActions from "../CardActions";
import {
  mutation,
  Data,
  Variables
} from "../../../apollo/Mutation/User/Password";
import { MessageContext } from "../../../context";
import Group from "../../../ui/Form/Group";
import { InputPassword } from "../../../ui/Form/Input";
import Label from "../../../ui/Form/Label";
import { getErrorMessage } from "../../../apollo/helpers/error";
import Title from "../Title";
import { PasswordHint } from "../../../ui/Form/Message/Hint";

const useStyles = makeStyles(({ spacing }) => ({
  form: {
    display: "flex",
    flexDirection: "column"
  },
  button: {
    marginTop: spacing(1),
    alignSelf: "flex-end"
  }
}));

export default function UserPassword(): JSX.Element {
  const setMessageInfo = useContext(MessageContext);
  const [oldPassword, setOldPassword] = useState("");
  const [newPassword, setNewPassword] = useState("");
  const classes = useStyles();
  const [mutate] = useMutation<Data, Variables>(mutation, {
    onCompleted: () => {
      setOldPassword("");
      setNewPassword("");
      setMessageInfo({ message: "We have changed your password." });
    },
    onError: error => setMessageInfo({ message: getErrorMessage(error) })
  });

  return (
    <Card>
      <form className={classes.form} onSubmit={event => event.preventDefault()}>
        <Title value="Password" />
        <CardContent>
          <Group>
            <Label htmlFor="old-password">Old password</Label>
            <InputPassword
              id="old-password"
              value={oldPassword}
              onChange={event => setOldPassword(event.target.value)}
            />
          </Group>
          <Group>
            <Label htmlFor="new-password">New password</Label>
            <InputPassword
              id="new-password"
              value={newPassword}
              aria-describedby="new-password-helper-text"
              onChange={event => setNewPassword(event.target.value)}
            />
            <PasswordHint id="new-password-helper-text" />
          </Group>
        </CardContent>
        <CardActions>
          <Button
            type="submit"
            variant="contained"
            color="primary"
            onClick={() =>
              mutate({ variables: { old: oldPassword, new: newPassword } })
            }
            className={classes.button}
          >
            Change
          </Button>
        </CardActions>
      </form>
    </Card>
  );
}
