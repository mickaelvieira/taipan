import { User } from "../types/users";

// @TODO to be implemented
export function isAdmin(user: User | null): boolean {
  if (!user) {
    return false;
  }
  return user.id === "1" ? true : false;
  // return false;
}
