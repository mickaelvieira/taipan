import React, { PropsWithChildren } from "react";
import AppLayout from "./App";
import OutLayout from "./Out";
import { User } from "../../types/users";
import useFirewall from "../../hooks/useFirewall";

interface Props {
  user: User | null;
}

export default function Layout({
  user,
  children,
}: PropsWithChildren<Props>): JSX.Element | null {
  const redirect = useFirewall(user);
  const Component = user ? AppLayout : OutLayout;
  if (redirect) {
    window.location.href = redirect;
    return null;
  }
  return <Component user={user}>{children}</Component>;
}
