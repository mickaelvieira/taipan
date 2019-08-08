import React, { PropsWithChildren } from "react";
import { useSubscription } from "@apollo/react-hooks";
import { User, UserEvent } from "../../types/users";
import { UserContext } from "../context";
import { userSubscription } from "../apollo/Subscription/User";

interface Props {
  loggedIn: User;
}

export default function AppUser({
  children,
  loggedIn
}: PropsWithChildren<Props>): JSX.Element {
  useSubscription<UserEvent>(userSubscription);
  return (
    <UserContext.Provider value={loggedIn}>{children}</UserContext.Provider>
  );
}
