import React, { useState } from "react";
import { makeStyles } from "@material-ui/core/styles";
import { RouteLoginProps } from "../../../types/routes";
import Button from "@material-ui/core/Button";
import Typography from "@material-ui/core/Typography";
import { login } from "../../../helpers/app";
import Link from "@material-ui/core/Link";
import { RouterLink } from "../../ui/Link";
import Group from "../../ui/Form/Group";
import Label from "../../ui/Form/Label";
import { InputBase, InputPassword } from "../../ui/Form/Input";

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
  },
  links: {
    display: "flex",
    justifyContent: "space-between"
  }
}));

export default function Signin(_: RouteLoginProps): JSX.Element {
  const classes = useStyles();
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  return (
    <div className={classes.container}>
      <Typography component="h2" variant="h5">
        Sign in
      </Typography>
      <form className={classes.form} onSubmit={event => event.preventDefault()}>
        <Group>
          <Label>Email address</Label>
          <InputBase
            id="username"
            value={username}
            onChange={event => setUsername(event.target.value)}
          />
        </Group>
        <Group>
          <Label>Password</Label>
          <InputPassword
            id="password"
            value={username}
            autoComplete="on"
            onChange={event => setPassword(event.target.value)}
          />
        </Group>
        <Button
          type="submit"
          variant="contained"
          color="primary"
          className={classes.button}
          onClick={() => {
            login(username, password)
              .then(() => {
                window.location.href = "/";
              })
              .catch(e => {
                console.warn(e);
              });
          }}
        >
          Login
        </Button>
      </form>
      <div className={classes.links}>
        <Link to="/join" component={RouterLink}>
          Create an account
        </Link>
        <Link to="/forgot-password" component={RouterLink}>
          Forgot password?
        </Link>
      </div>
    </div>
  );
}
