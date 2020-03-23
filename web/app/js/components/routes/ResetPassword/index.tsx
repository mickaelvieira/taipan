import React, { useState } from "react";
import { makeStyles } from "@material-ui/core/styles";
import { RouteForgotPasswordProps } from "../../../types/routes";
import Button from "@material-ui/core/Button";
import Link from "@material-ui/core/Link";
import Typography from "@material-ui/core/Typography";
import { resetPassword } from "../../../helpers/app";
import { RouterLink } from "../../ui/Link";
import Group from "../../ui/Form/Group";
import Label from "../../ui/Form/Label";
import { PasswordHint } from "../../ui/Form/Message/Hint";
import { ErrorMessage, SuccessMessage } from "../../ui/Form/Message";
import { InputPassword } from "../../ui/Form/Input";

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
  message: {
    textAlign: "center",
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
    justifyContent: "center",
    marginTop: spacing(2),
  },
}));

export default function ResetPassword(
  _: RouteForgotPasswordProps
): JSX.Element {
  const classes = useStyles();
  const [error, setError] = useState("");
  const [message, setMessage] = useState("");
  const [password, setPassword] = useState("");
  const url = new URL(`${document.location}`);
  const params = url.searchParams;
  const token = params.get("token");

  return (
    <div className={classes.container}>
      {message && (
        <>
          <SuccessMessage className={classes.message}>{message}</SuccessMessage>
          <div className={classes.links}>
            <Link to="/signin" component={RouterLink}>
              Sign in
            </Link>
          </div>
        </>
      )}

      {!message && (
        <>
          <Typography component="h2" variant="h5">
            Change your password
          </Typography>
          <form
            className={classes.form}
            onSubmit={(event) => event.preventDefault()}
          >
            <Group>
              <Label htmlFor="password">Enter your new your password</Label>
              <InputPassword
                id="password"
                value={password}
                aria-describedby="password-helper-text"
                onChange={(event) => setPassword(event.target.value)}
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
                setMessage("");
                setError("");
                resetPassword(token, password)
                  .then(({ error }) => {
                    if (error) {
                      setError(error.error);
                    } else {
                      setMessage(
                        `We have reset your password, you can now log into your account.`
                      );
                    }
                  })
                  .catch((e) => {
                    setError(e.message);
                  });
              }}
            >
              Change my password
            </Button>
          </form>
        </>
      )}
    </div>
  );
}
