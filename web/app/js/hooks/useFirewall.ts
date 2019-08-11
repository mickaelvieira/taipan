import { User } from "../types/users";

const home = "/";
const login = "/sign-in";
const unsecured = [login, "/join", "/sign-out"];

export default function useFirewall(user: User | null): string | null {
  const url = new URL(`${document.location}`);

  if (!user && !unsecured.includes(url.pathname)) {
    return login;
  }
  if (user && location.pathname === login) {
    return home;
  }

  return null;
}
