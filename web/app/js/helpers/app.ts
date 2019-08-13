import { User } from "../types/users";
import { APIResponse } from "../types/app";

/* eslint @typescript-eslint/no-explicit-any: "off" */

export function getAppInfoFromEnv(): [string, string] {
  return [process.env.APP_NAME || "", process.env.APP_VERSION || ""];
}

function processResponse<T>(result: any): APIResponse<T> {
  if (typeof result.error !== "undefined") {
    return { error: result };
  }
  return { result };
}

function getJSONRequest(endpoint: string, body: any): Request {
  return new Request(endpoint, {
    method: "POST",
    headers: {
      "Content-Type": "application/json; charset=utf-8"
    },
    body: JSON.stringify(body)
  });
}

export function join(
  email: string,
  password: string
): Promise<APIResponse<User>> {
  return fetch(
    getJSONRequest("/sign-up", {
      email,
      password
    })
  )
    .then(response => response.json())
    .then(json => processResponse<User>(json));
}

export function login(
  username: string,
  password: string
): Promise<APIResponse<User>> {
  return fetch(
    getJSONRequest("/login", {
      username,
      password
    })
  )
    .then(response => response.json())
    .then(json => processResponse<User>(json));
}

export function logout(): Promise<APIResponse<{}>> {
  return fetch(getJSONRequest("/logout", {}))
    .then(response => response.json())
    .then(json => processResponse<User>(json));
}
