import React, { useContext } from "react";
import { makeStyles } from "@material-ui/core/styles";
import List from "@material-ui/core/List";
import Card from "@material-ui/core/Card";
import CardContent from "@material-ui/core/CardContent";
import { MessageContext } from "../../../context";
import Title from "../Title";
import UserEmail from "./Email";
import FormUserEmail from "./Form";
import { User } from "../../../../types/users";

const useStyles = makeStyles(() => ({
  card: {
    marginTop: 16
  }
}));

interface Props {
  user: User;
}

export default function UserEmails({ user }: Props): JSX.Element {
  const setMessageInfo = useContext(MessageContext);
  const classes = useStyles();
  let canAdd = true;

  if (user.emails.length === 1 && !user.emails[0].isConfirmed) {
    canAdd = false;
  }

  return (
    <Card className={classes.card}>
      <Title value="Emails" />
      <CardContent>
        <List>
          {user.emails.map(email => (
            <UserEmail
              key={email.id}
              email={email}
              onDeleted={() => {
                setMessageInfo({ message: "Email was deleted" });
              }}
            />
          ))}
        </List>
        {canAdd && (
          <FormUserEmail
            onCreated={() => {
              setMessageInfo({ message: "Email was created" });
            }}
          />
        )}
      </CardContent>
    </Card>
  );
}
