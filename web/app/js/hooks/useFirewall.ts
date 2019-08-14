import { User } from "../types/users";

const home = "/";
const login = "/sign-in";
const unsecured = [login, "/join", "/sign-out"];

function shouldRedirect(redirect: string): boolean {
  return location.pathname !== redirect;
}

export default function useFirewall(user: User | null): string | null {
  const url = new URL(`${document.location}`);

  if (!user && !unsecured.includes(url.pathname)) {
    return shouldRedirect(login) ? login : null;
  }
  if (user && location.pathname === login) {
    return shouldRedirect(home) ? home : null;
  }

  return null;
}
