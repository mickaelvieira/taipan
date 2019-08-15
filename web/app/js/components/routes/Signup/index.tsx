import React, { useState } from "react";
import { makeStyles } from "@material-ui/core/styles";
import { RouteLoginProps } from "../../../types/routes";
import Button from "@material-ui/core/Button";
import Link from "@material-ui/core/Link";
import Typography from "@material-ui/core/Typography";
import { join } from "../../../helpers/app";
import { RouterLink } from "../../ui/Link";
import Group from "../../ui/Form/Group";
import Label from "../../ui/Form/Label";
import { PasswordHint, EmailHint } from "../../ui/Form/Message/Hint";
import { ErrorMessage } from "../../ui/Form/Message";
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
    justifyContent: "center"
  }
}));

export default function Signup(_: RouteLoginProps): JSX.Element {
  const classes = useStyles();
  const [error, setError] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  return (
    <div className={classes.container}>
      <Typography component="h2" variant="h5">
        Create your personal account
      </Typography>
      <form className={classes.form} onSubmit={event => event.preventDefault()}>
        <Group>
          <Label htmlFor="email">Email address</Label>
          <InputBase
            id="email"
            value={email}
            aria-describedby="email-helper-text"
            onChange={event => setEmail(event.target.value)}
          />
          <EmailHint id="email-helper-text" />
        </Group>
        <Group>
          <Label htmlFor="password">Password</Label>
          <InputPassword
            id="password"
            value={password}
            autoComplete="on"
            aria-describedby="password-helper-text"
            onChange={event => setPassword(event.target.value)}
          />
          <PasswordHint id="password-helper-text" />
        </Group>
        {error && <ErrorMessage>{error}</ErrorMessage>}
        <Button
          type="submit"
          variant="contained"
          color="primary"
          className={classes.button}
          onClick={() => {
            join(email, password)
              .then(({ error, result }) => {
                if (error) {
                  setError(error.error);
                } else if (result) {
                  window.location.href = "/";
                }
              })
              .catch(e => {
                setError(e.message);
              });
          }}
        >
          Create an account
        </Button>
      </form>
      <div className={classes.links}>
        <Link to="/sign-in" component={RouterLink}>
          I already have an account
        </Link>
      </div>
    </div>
  );
}
