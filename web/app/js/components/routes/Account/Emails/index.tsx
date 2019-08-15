import React, { useContext } from "react";
import List from "@material-ui/core/List";
import CardContent from "@material-ui/core/CardContent";
import { MessageContext } from "../../../context";
import Title from "../Title";
import Card from "../Card";
import UserEmail from "./Email";
import FormUserEmail from "./Form";
import { User } from "../../../../types/users";

interface Props {
  user: User;
}

export default function UserEmails({ user }: Props): JSX.Element {
  const setMessageInfo = useContext(MessageContext);
  let canAdd = true;

  // If the user has only one email and it hasn't been confirmed yet
  // the user is not allowed to add a new one.
  if (user.emails.length === 1 && !user.emails[0].isConfirmed) {
    canAdd = false;
  }

  return (
    <Card>
      <Title value="Emails" />
      <CardContent>
        <List>
          {user.emails.map(email => (
            <UserEmail
              key={email.id}
              email={email}
              onDeletionFailure={message => {
                setMessageInfo({ message });
              }}
              onStatusFailure={message => {
                setMessageInfo({ message });
              }}
            />
          ))}
        </List>
        {canAdd && (
          <FormUserEmail
            onCreationFailure={message => {
              setMessageInfo({ message });
            }}
          />
        )}
      </CardContent>
    </Card>
  );
}
