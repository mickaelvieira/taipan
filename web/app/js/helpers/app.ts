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

export async function join(
  email: string,
  password: string
): Promise<APIResponse<User>> {
  const response = await fetch(
    getJSONRequest("/join", {
      email,
      password
    })
  );
  const json = await response.json();
  return processResponse<User>(json);
}

export async function login(
  email: string,
  password: string
): Promise<APIResponse<User>> {
  const response = await fetch(
    getJSONRequest("/signin", {
      email,
      password
    })
  );
  const json = await response.json();
  return processResponse<User>(json);
}

export async function logout(): Promise<APIResponse<{}>> {
  const response = await fetch(getJSONRequest("/signout", {}));
  const json = await response.json();
  return processResponse<User>(json);
}

export async function askForResetEmail(
  email: string
): Promise<APIResponse<{}>> {
  const response = await fetch(
    getJSONRequest("/forgot-password", {
      email
    })
  );
  const json = await response.json();
  return processResponse<{}>(json);
}

export async function resetPassword(
  token: string | null,
  password: string
): Promise<APIResponse<{}>> {
  const response = await fetch(
    getJSONRequest("/reset-password", {
      token,
      password
    })
  );
  const json = await response.json();
  return processResponse<{}>(json);
}

export async function confirmEmail(
  token: string | null
): Promise<APIResponse<{}>> {
  const response = await fetch(
    getJSONRequest("/confirm-email", {
      token
    })
  );
  const json = await response.json();
  return processResponse<{}>(json);
}
