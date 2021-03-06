import React, { useState } from "react";
import { makeStyles } from "@material-ui/core/styles";
import { RouteSigninProps } from "../../../types/routes";
import Button from "@material-ui/core/Button";
import Typography from "@material-ui/core/Typography";
import { login } from "../../../helpers/app";
import Link from "@material-ui/core/Link";
import { RouterLink } from "../../ui/Link";
import Group from "../../ui/Form/Group";
import Label from "../../ui/Form/Label";
import { ErrorMessage } from "../../ui/Form/Message";
import { InputBase, InputPassword } from "../../ui/Form/Input";

const useStyles = makeStyles(({ spacing, palette, breakpoints }) => ({
  container: {
    width: "300px",
    alignSelf: "center",
    border: `1px solid ${palette.grey[500]}`,
    padding: spacing(2),
    [breakpoints.up("sm")]: {
      width: "600px",
    },
  },
  form: {
    display: "flex",
    flexDirection: "column",
  },
  button: {
    margin: `${spacing(2)}px 0`,
  },
  links: {
    display: "flex",
    justifyContent: "space-between",
  },
}));

export default function Signin(_: RouteSigninProps): JSX.Element {
  const classes = useStyles();
  const [error, setError] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  return (
    <div className={classes.container}>
      <Typography component="h2" variant="h5">
        Sign in
      </Typography>
      <form
        className={classes.form}
        onSubmit={(event) => event.preventDefault()}
      >
        <Group>
          <Label htmlFor="email">Email address</Label>
          <InputBase
            id="email"
            type="email"
            value={email}
            onChange={(event) => setEmail(event.target.value)}
          />
        </Group>
        <Group>
          <Label htmlFor="password">Password</Label>
          <InputPassword
            id="password"
            value={password}
            onChange={(event) => setPassword(event.target.value)}
          />
        </Group>
        {error && <ErrorMessage>{error}</ErrorMessage>}
        <Button
          type="submit"
          variant="contained"
          color="primary"
          className={classes.button}
          onClick={() => {
            login(email, password)
              .then(({ error, result }) => {
                if (error) {
                  setError(error.error);
                } else if (result) {
                  window.location.href = "/";
                }
              })
              .catch((e) => {
                setError(e.message);
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
