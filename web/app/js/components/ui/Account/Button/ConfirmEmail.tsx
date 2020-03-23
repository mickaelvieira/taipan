import React, { useState } from "react";
import { makeStyles } from "@material-ui/core/styles";
import { useMutation } from "@apollo/react-hooks";
import {
  mutation,
  Variables,
  Data,
} from "../../../apollo/Mutation/User/ConfirmEmail";
import { getErrorMessage } from "../../../apollo/helpers/error";
import { Email } from "../../../../types/users";
import { Button } from "@material-ui/core";

const useStyles = makeStyles(({ palette }) => ({
  message: {
    color: palette.common.black,
    padding: "6px 8px",
  },
}));

interface Props {
  email: Email;
  showConfirm?: boolean;
  className?: string;
  onSuccess: (message: string) => void;
  onFailure: (message: string) => void;
}

export default function ConfirmEmail({
  email,
  showConfirm,
  className,
  onFailure,
  onSuccess,
}: Props): JSX.Element | null {
  const classes = useStyles();
  const [sent, setIsSent] = useState(false);
  const [send, { loading }] = useMutation<Data, Variables>(mutation, {
    onError: (error) => onFailure(getErrorMessage(error)),
    onCompleted: () => {
      setIsSent(true);
      onSuccess(`We have sent you an email to ${email.value}`);
    },
  });

  if (sent) {
    return !showConfirm ? null : (
      <span className={`${classes.message} ${className ? className : ""}`}>
        Email was sent to {email.value}{" "}
      </span>
    );
  }

  return (
    <Button
      className={`${className ? className : ""}`}
      disabled={loading}
      onClick={() =>
        send({
          variables: { email: email.value },
        })
      }
    >
      Resend
    </Button>
  );
}
