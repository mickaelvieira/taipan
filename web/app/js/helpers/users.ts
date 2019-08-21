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

export function getFullname(user: User | null): string {
  if (!user) {
    return "";
  }

  const names = [];
  if (user.firstname) {
    names.push(user.firstname);
  }
  if (user.lastname) {
    names.push(user.lastname);
  }

  return names.join(" ");
}

export function getEmailHandle(email: Email | null): string {
  if (!email) {
    return "";
  }

  const parts = email.value.split("@");
  if (parts.length < 2) {
    return "";
  }

  return parts[0];
}
