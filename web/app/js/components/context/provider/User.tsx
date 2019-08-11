import React, { PropsWithChildren } from "react";
import { useSubscription } from "@apollo/react-hooks";
import { User, UserEvent } from "../../../types/users";
import { UserContext } from "..";
import { userSubscription } from "../../apollo/Subscription/User";

interface Props {
  user: User;
}

export default function UserContextProvider({
  children,
  user
}: PropsWithChildren<Props>): JSX.Element {
  useSubscription<UserEvent>(userSubscription);
  return <UserContext.Provider value={user}>{children}</UserContext.Provider>;
}
