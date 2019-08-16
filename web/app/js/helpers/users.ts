import { User, Email } from "../types/users";

// @TODO to be implemented
export function isAdmin(user: User | null): boolean {
  if (!user) {
    return false;
  }
  return user.id === "1" ? true : false;
  // return false;
}

export function getPrimaryEmail(user: User | null): Email | null {
  if (!user) {
    return null;
  }

  const primary = user.emails.find(({ isPrimary }) => isPrimary);
  return primary ? primary : null;
}
