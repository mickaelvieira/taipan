import React, { useState } from "react";
import { useMutation } from "@apollo/react-hooks";
import { makeStyles } from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import TextField from "@material-ui/core/TextField";
import { Subscription } from "../../../../../types/subscription";
import {
  mutation,
  Data,
  Variables
} from "../../../../apollo/Mutation/Subscriptions/Subscription";

const useStyles = makeStyles(() => ({
  form: {
    display: "flex",
    flexDirection: "column"
  },
  button: {
    alignSelf: "flex-end"
  }
}));

interface Props {
  onSubscriptionCreated: (subscription: Subscription) => void;
}

export default function FormSubscription({
  onSubscriptionCreated
}: Props): JSX.Element {
  const classes = useStyles();
  const [url, setUrl] = useState("");
  const [mutate, { loading, error }] = useMutation<Data, Variables>(mutation, {
    onCompleted: ({ subscriptions: { subscription } }) => {
      onSubscriptionCreated(subscription);
    }
  });

  return (
    <form className={classes.form}>
      <TextField
        autoFocus
        margin="dense"
        id="feed_url"
        label="URL"
        placeholder="https://"
        type="url"
        value={url}
        error={!!error}
        autoComplete="off"
        autoCapitalize="off"
        autoCorrect="off"
        helperText={!error ? "" : error.message}
        onChange={event => setUrl(event.target.value)}
        fullWidth
      />
      <Button
        onClick={() =>
          mutate({
            variables: { url }
          })
        }
        className={classes.button}
        color="primary"
        disabled={loading}
      >
        Add
      </Button>
    </form>
  );
}
