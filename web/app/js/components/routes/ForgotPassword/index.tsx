import React, { useState } from "react";
import { makeStyles } from "@material-ui/core/styles";
import { RouteForgotPasswordProps } from "../../../types/routes";
import Button from "@material-ui/core/Button";
import Link from "@material-ui/core/Link";
import Typography from "@material-ui/core/Typography";
import { askForResetEmail } from "../../../helpers/app";
import { RouterLink } from "../../ui/Link";
import Group from "../../ui/Form/Group";
import Label from "../../ui/Form/Label";
import { ErrorMessage, SuccessMessage } from "../../ui/Form/Message";
import { InputBase } from "../../ui/Form/Input";

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
  message: {
    textAlign: "center",
    marginBottom: spacing(2)
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

export default function ForgotPassword(
  _: RouteForgotPasswordProps
): JSX.Element {
  const classes = useStyles();
  const [error, setError] = useState("");
  const [message, setMessage] = useState("");
  const [email, setEmail] = useState("");

  return (
    <div className={classes.container}>
      {message && (
        <SuccessMessage className={classes.message}>{message}</SuccessMessage>
      )}
      {!message && (
        <>
          <Typography component="h2" variant="h5">
            Reset your password
          </Typography>
          <form
            className={classes.form}
            onSubmit={event => event.preventDefault()}
          >
            <Group>
              <Label htmlFor="email">
                Enter your email address and we will send you a link to reset
                your password.
              </Label>
              <InputBase
                id="email"
                value={email}
                aria-describedby="email-helper-text"
                onChange={event => setEmail(event.target.value)}
              />
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
                askForResetEmail(email)
                  .then(({ error }) => {
                    if (error) {
                      setError(error.error);
                    } else {
                      setMessage(
                        `If you have an account registered with us, we have sent you an email to ${email} to reset your password`
                      );
                    }
                  })
                  .catch(e => {
                    setError(e.message);
                  });
              }}
            >
              Send me an email to reset my password
            </Button>
          </form>
          <div className={classes.links}>
            <Link to="/signin" component={RouterLink}>
              Actually I remember my password
            </Link>
          </div>
        </>
      )}
    </div>
  );
}
