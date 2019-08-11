import React, { useState } from "react";
import { useMutation } from "@apollo/react-hooks";
import { makeStyles } from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import { Subscription } from "../../../../../types/subscription";
import {
  mutation,
  Data,
  Variables
} from "../../../../apollo/Mutation/Subscriptions/Subscription";
import Group from "../../../../ui/Form/Group";
import Label from "../../../../ui/Form/Label";
import Hint from "../../../../ui/Form/Hint";
import { InputBase } from "../../../../ui/Form/Input";

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
      <Group>
        <Label htmlFor="feed-url">URL</Label>
        <InputBase
          id="feed-url"
          autoFocus
          placeholder="https://"
          type="url"
          value={url}
          error={!!error}
          autoComplete="off"
          autoCapitalize="off"
          autoCorrect="off"
          onChange={event => setUrl(event.target.value)}
        />
        {error && <Hint>{error.message}</Hint>}
      </Group>
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
