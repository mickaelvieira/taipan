import React, { PropsWithChildren, useState } from "react";
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
  const [user, setUser] = useState(loggedIn);
  return (
    <UserContext.Provider value={user}>
      <UserSubscription update={setUser} />
      {children}
    </UserContext.Provider>
  );
}
