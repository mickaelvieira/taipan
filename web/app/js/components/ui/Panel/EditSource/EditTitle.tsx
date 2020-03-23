import React, { useState } from "react";
import { useMutation } from "@apollo/react-hooks";
import { makeStyles } from "@material-ui/core/styles";
import IconButton from "@material-ui/core/IconButton";
import EditIcon from "@material-ui/icons/Edit";
import CancelIcon from "@material-ui/icons/Cancel";
import InputBase from "@material-ui/core/InputBase";
import {
  mutation,
  Data,
  Variables,
} from "../../../apollo/Mutation/Syndication/Title";
import Title from "../../../routes/Syndication/Title";
import { Source } from "../../../../types/syndication";

const useStyles = makeStyles(({ palette }) => ({
  editor: {
    display: "flex",
    alignItems: "center",
  },
  title: {
    width: "100%",
  },
  button: {
    paddingTop: 0,
    paddingBottom: 0,
  },
  input: {
    border: `1px solid ${palette.grey[200]}`,
  },
}));

interface Props {
  source: Source;
}

export default function EditTitle({ source }: Props): JSX.Element {
  const classes = useStyles();
  const [value, setValue] = useState(source.title);
  const [editMode, setEditMode] = useState(false);
  const [mutate] = useMutation<Data, Variables>(mutation, {
    onCompleted: () => setEditMode(false),
  });

  return (
    <>
      {!editMode && (
        <div className={classes.editor}>
          <Title item={source} className={classes.title} />
          <IconButton
            onClick={() => setEditMode(true)}
            className={classes.button}
          >
            <EditIcon fontSize="small" />
          </IconButton>
        </div>
      )}
      {editMode && (
        <form
          className={classes.editor}
          onSubmit={(event) => event.preventDefault()}
        >
          <InputBase
            fullWidth
            className={classes.input}
            autoFocus
            value={value}
            onChange={(event) => setValue(event.target.value)}
          />
          <IconButton
            type="submit"
            className={classes.button}
            onClick={() =>
              mutate({
                variables: { url: source.url, title: value },
              })
            }
          >
            <EditIcon fontSize="small" />
          </IconButton>
          <IconButton
            onClick={() => setEditMode(false)}
            className={classes.button}
          >
            <CancelIcon fontSize="small" />
          </IconButton>
        </form>
      )}
    </>
  );
}
