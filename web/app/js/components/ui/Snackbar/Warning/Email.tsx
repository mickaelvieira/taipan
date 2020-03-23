import React, { PropsWithChildren } from "react";
import { SnackbarWarning } from "../";
import { User } from "../../../../types/users";
import { getPrimaryEmail } from "../../../../helpers/users";
import { ConfirmEmailButton } from "../../Account/Button";

interface Props {
  user: User | null;
  onRendSuccess: (message: string) => void;
  onRendFailure: (message: string) => void;
}

export default function SnackbarEmailWarning({
  user,
  onRendSuccess,
  onRendFailure,
}: PropsWithChildren<Props>): JSX.Element | null {
  const email = getPrimaryEmail(user);
  if (!email) {
    return null;
  }

  return (
    <SnackbarWarning
      open={!!(email && !email.isConfirmed)}
      message="Your primary email address has not been confirm. We have sent you an email to valid it."
      action={[
        <ConfirmEmailButton
          key={0}
          email={email}
          onSuccess={onRendSuccess}
          onFailure={onRendFailure}
        />,
      ]}
    />
  );
}
