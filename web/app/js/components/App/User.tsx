import React, { PropsWithChildren } from "react";
import { User } from "../../types/users";
import { UserContext } from "../context";
import UserSubscription from "../apollo/Subscription/User";

interface Props {
  loggedIn: User;
}

export default function AppUser({
  children,
  loggedIn
}: PropsWithChildren<Props>): JSX.Element {
  return (
    <UserContext.Provider value={loggedIn}>
      <UserSubscription />
      {children}
    </UserContext.Provider>
  );
}
