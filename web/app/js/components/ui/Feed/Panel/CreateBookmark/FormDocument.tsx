import React, { useState } from "react";
import { useMutation } from "@apollo/react-hooks";
import { makeStyles } from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import TextField from "@material-ui/core/TextField";
import Typography from "@material-ui/core/Typography";
import {
  mutation,
  Data,
  Variables
} from "../../../../apollo/Mutation/Documents/Create";
import { Document } from "../../../../../types/document";

const useStyles = makeStyles(() => ({
  form: {
    display: "flex",
    flexDirection: "column"
  },
  button: {
    marginTop: 16,
    alignSelf: "flex-end"
  }
}));

interface Props {
  onFetchDocument: (document: Document) => void;
}

export default function FormDocument({ onFetchDocument }: Props): JSX.Element {
  const classes = useStyles();
  const [url, setUrl] = useState("");
  const [mutate, { loading, error }] = useMutation<Data, Variables>(mutation, {
    onCompleted: ({ documents: { create } }) => onFetchDocument(create)
  });

  return (
    <form className={classes.form}>
      <Typography paragraph>
        Enter the URL of the page you would like to bookmark
      </Typography>
      <TextField
        autoFocus
        margin="dense"
        id="bookmark_url"
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
        color="primary"
        disabled={loading}
        className={classes.button}
      >
        Preview
      </Button>
    </form>
  );
}
