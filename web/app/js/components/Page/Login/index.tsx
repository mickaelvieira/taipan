import React, { useState } from "react";
import { makeStyles } from "@material-ui/core/styles";
import { RouteLoginProps } from "../../../types/routes";
import Button from "@material-ui/core/Button";
import TextField from "@material-ui/core/TextField";
import { login } from "../../../helpers/app";

const useStyles = makeStyles(({ spacing, palette, breakpoints }) => ({
  container: {
    width: "300px",
    alignSelf: "center",
    border: `1px solid ${palette.grey[500]}`,
    padding: spacing(2),
    [breakpoints.up("sm")]: {
      width: "600px"
    }
  },
  form: {
    display: "flex",
    flexDirection: "column"
  },
  button: {
    margin: `${spacing(2)}px 0`
  }
}));

export default function Login(_: RouteLoginProps): JSX.Element {
  const classes = useStyles();
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  return (
    <div className={classes.container}>
      <form className={classes.form} onSubmit={event => event.preventDefault()}>
        <TextField
          fullWidth
          margin="normal"
          id="username"
          label="Username"
          value={username}
          onChange={event => setUsername(event.target.value)}
        />
        <TextField
          fullWidth
          margin="normal"
          id="password"
          label="Password"
          value={password}
          type="password"
          onChange={event => setPassword(event.target.value)}
        />
        <Button
          type="submit"
          variant="contained"
          color="primary"
          className={classes.button}
          onClick={() => {
            login(username, password)
              .then(user => {
                console.log(user);
                window.location = "/";
              })
              .catch(e => {
                console.log(e);
              });
          }}
        >
          Login
        </Button>
      </form>
    </div>
  );
}
