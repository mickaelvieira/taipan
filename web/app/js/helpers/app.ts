import { User } from "../types/users";

export async function login(username: string, password: string): Promise<User> {
  try {
    const result = await fetch("/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({
        username,
        password
      })
    }).then(response => response.json());
    return result;
  } catch (e) {
    throw e;
  }
}

export async function logout(): Promise<string> {
  try {
    const result = await fetch("/logout", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({})
    }).then(response => response.text());
    return result;
  } catch (e) {
    throw e;
  }
}
