import React, { useState } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import TextField from "@material-ui/core/TextField";
import CreateUserEmailMutation from "../../../apollo/Mutation/User/CreateEmail";

const useStyles = makeStyles(() => ({
  form: {
    display: "flex",
    flexDirection: "column",
    alignItems: "flex-start"
  },
  button: {
    alignSelf: "flex-end"
  }
}));

interface Props {
  onCreated: () => void;
}

export default function UserEmailForm({ onCreated }: Props): JSX.Element {
  const [email, setEmail] = useState("");
  const classes = useStyles();

  return (
    <CreateUserEmailMutation
      onCompleted={() => {
        setEmail("");
        onCreated();
      }}
    >
      {mutate => (
        <form
          className={classes.form}
          onSubmit={event => event.preventDefault()}
        >
          <TextField
            fullWidth
            margin="normal"
            id="email"
            label="Email"
            value={email}
            onChange={event => setEmail(event.target.value)}
          />
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
      )}
    </CreateUserEmailMutation>
  );
}
