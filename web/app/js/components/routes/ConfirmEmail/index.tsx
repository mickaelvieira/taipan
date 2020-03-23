import React, { useState, useEffect, useContext } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Typography from "@material-ui/core/Typography";
import Link from "@material-ui/core/Link";
import { confirmEmail } from "../../../helpers/app";
import { ErrorMessage, SuccessMessage } from "../../ui/Form/Message";
import Loader from "../../ui/Loader";
import { RouteConfirmEmailProps } from "../../../types/routes";
import ReloadUser from "./ReloadUser";
import { UserContext } from "../../context";

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
    marginBottom: spacing(2),
  },
  links: {
    display: "flex",
    justifyContent: "center",
    marginTop: spacing(2),
  },
}));

export default function ConfirmEmail(_: RouteConfirmEmailProps): JSX.Element {
  const classes = useStyles();
  const user = useContext(UserContext);
  const [error, setError] = useState("");
  const [message, setMessage] = useState("");
  const [isLoading, setIsLoading] = useState(true);
  const url = new URL(`${document.location}`);
  const params = url.searchParams;
  const token = params.get("token");

  useEffect(() => {
    confirmEmail(token)
      .then(({ error }) => {
        if (error) {
          setIsLoading(false);
          setError(error.error);
        } else {
          setIsLoading(false);
          setMessage(`Many thanks. Your email address has been confirmed.`);
        }
      })
      .catch((e) => {
        setError(e.message);
      });
  }, [token]);

  return (
    <div className={classes.container}>
      {message && (
        <>
          <SuccessMessage className={classes.message}>{message}</SuccessMessage>
          <div className={classes.links}>
            <Link href="/account">Go to my account</Link>
            {/*
              The user could be logged in and since we don't use apollo to confirm the email address
              we can't update the cache. We then need to refresh the cache in the background
              otherwise the user's email will be marked as unconfirmed.
            */}
            {user && <ReloadUser />}
          </div>
        </>
      )}

      {error && (
        <>
          <ErrorMessage className={classes.message}>{error}</ErrorMessage>
          <Typography component="p" className={classes.message}>
            We can send you an new email to confirm your address. Go to your
            account and press &quot;resend&quot; button next to the email you
            want to confirm.
          </Typography>
          <div className={classes.links}>
            <Link href="/account">Go to my account</Link>
          </div>
        </>
      )}

      {isLoading && <Loader />}
    </div>
  );
}
