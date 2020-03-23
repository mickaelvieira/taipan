import React, { useState } from "react";
import { useMutation } from "@apollo/react-hooks";
import { makeStyles } from "@material-ui/core/styles";
import IconButton from "@material-ui/core/IconButton";
import EditIcon from "@material-ui/icons/Edit";
import InputBase from "@material-ui/core/InputBase";
import {
  mutation,
  Data,
  Variables,
} from "../../../apollo/Mutation/Syndication/Tags/Create";
import { query } from "../../../apollo/Query/Tags";

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

export default function FormTag(): JSX.Element {
  const classes = useStyles();
  const [value, setValue] = useState("");
  const [mutate] = useMutation<Data, Variables>(mutation, {
    onCompleted: () => setValue(""),
    refetchQueries: [
      {
        query,
      },
    ],
  });

  return (
    <>
      <form
        className={classes.editor}
        onSubmit={(event) => event.preventDefault()}
      >
        <InputBase
          fullWidth
          className={classes.input}
          value={value}
          onChange={(event) => setValue(event.target.value)}
        />
        <IconButton
          type="submit"
          className={classes.button}
          onClick={() =>
            mutate({
              variables: { label: value },
            })
          }
        >
          <EditIcon fontSize="small" />
        </IconButton>
      </form>
    </>
  );
}
