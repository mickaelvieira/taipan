import React, { PropsWithChildren } from "react";
import { withRouter, Redirect } from "react-router";
import AppLayout from "./App";
import OutLayout from "./Out";
import { RoutesProps } from "../../types/routes";
import { User } from "../../types/users";

interface Props extends RoutesProps {
  user: User | null;
}

const home = "/";
const login = "/sign-in";
const unsecured = [login, "/sign-up", "/sign-out"];

export default withRouter(function Layout({
  user,
  children,
  location
}: PropsWithChildren<Props>) {
  const Component = user ? AppLayout : OutLayout;
  if (!user && !unsecured.includes(location.pathname)) {
    return <Redirect exact strict to={login} />;
  }
  if (user && location.pathname === login) {
    return <Redirect exact strict to={home} />;
  }
  return <Component user={user}>{children}</Component>;
});
