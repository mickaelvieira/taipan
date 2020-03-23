import React, { PropsWithChildren } from "react";
// import { useSubscription } from "@apollo/react-hooks";
// import { User, UserEvent } from "../../../types/users";
import { User } from "../../../types/users";
import { UserContext } from "..";
// import { userSubscription } from "../../apollo/Subscription/User";

interface Props {
  user: User;
}

export default function UserContextProvider({
  children,
  user,
}: PropsWithChildren<Props>): JSX.Element {
  // useSubscription<UserEvent>(userSubscription, {
  //   onSubscriptionData: ({ subscriptionData }) => {
  //     console.log(subscriptionData);
  //   }
  // });
  return <UserContext.Provider value={user}>{children}</UserContext.Provider>;
}
