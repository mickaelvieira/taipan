import React, { useState } from "react";
import { useMutation } from "@apollo/react-hooks";
import { makeStyles } from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import {
  mutation,
  Data,
  Variables
} from "../../../apollo/Mutation/User/CreateEmail";
import Group from "../../../ui/Form/Group";
import { InputBase } from "../../../ui/Form/Input";
import Label from "../../../ui/Form/Label";

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

interface Props {
  onCreated: () => void;
}

export default function UserEmailForm({ onCreated }: Props): JSX.Element {
  const [email, setEmail] = useState("");
  const classes = useStyles();
  const [mutate] = useMutation<Data, Variables>(mutation, {
    onCompleted: () => {
      setEmail("");
      onCreated();
    }
  });

  return (
    <form className={classes.form} onSubmit={event => event.preventDefault()}>
      <Group>
        <Label htmlFor="new-email">Email</Label>
        <InputBase
          id="new-email"
          value={email}
          onChange={event => setEmail(event.target.value)}
        />
      </Group>
      <Button
        type="submit"
        variant="contained"
        color="primary"
        onClick={() => mutate({ variables: { email } })}
        className={classes.button}
      >
        Add email
      </Button>
    </form>
  );
}
