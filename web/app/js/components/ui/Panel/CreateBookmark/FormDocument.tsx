import React, { useState } from "react";
import { useMutation } from "@apollo/react-hooks";
import { makeStyles } from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import Typography from "@material-ui/core/Typography";
import {
  mutation,
  Data,
  Variables
} from "../../../apollo/Mutation/Documents/Create";
import { Document } from "../../../../types/document";
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
    marginTop: 16,
    alignSelf: "flex-end"
  }
}));

interface Props {
  onDocumentFetched: (document: Document) => void;
}

export default function FormDocument({
  onDocumentFetched
}: Props): JSX.Element {
  const classes = useStyles();
  const [url, setUrl] = useState("");
  const [createDocument, { loading, error }] = useMutation<Data, Variables>(
    mutation,
    {
      onCompleted: ({ documents: { create } }) => onDocumentFetched(create)
    }
  );

  return (
    <form className={classes.form}>
      <Typography paragraph>
        Enter the URL of the page you would like to bookmark
      </Typography>
      <Group>
        <Label htmlFor="bookmark-url">URL</Label>
        <InputBase
          id="bookmark-url"
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
          createDocument({
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
