import React, { useState } from "react";
import { useMutation } from "@apollo/react-hooks";
import { makeStyles } from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import { Subscription } from "../../../../types/subscription";
import {
  mutation,
  Data,
  Variables
} from "../../../apollo/Mutation/Subscriptions/Subscription";
import Group from "../../Form/Group";
import Label from "../../Form/Label";
import { ErrorMessage } from "../../Form/Message";
import { InputBase } from "../../Form/Input";

const useStyles = makeStyles(({ palette }) => ({
  form: {
    display: "flex",
    flexDirection: "column"
  },
  input: {
    borderRadius: 0,
    border: 0,
    borderBottom: `1px solid ${palette.grey[400]}`,
    paddingRight: 0,
    paddingLeft: 0
  },
  button: {
    alignSelf: "flex-end"
  }
}));

interface Props {
  onCreated: (subscription: Subscription) => void;
}

export default function FormSubscription({ onCreated }: Props): JSX.Element {
  const classes = useStyles();
  const [url, setUrl] = useState("");
  const [mutate, { loading, error }] = useMutation<Data, Variables>(mutation, {
    onCompleted: ({ subscriptions: { subscription } }) => {
      onCreated(subscription);
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
          className={classes.input}
        />
        {error && <ErrorMessage>{error.message}</ErrorMessage>}
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
