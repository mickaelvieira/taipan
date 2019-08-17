export interface AppInfo {
  name: string;
  version: string;
}

export interface APIError {
  error: string;
}

export interface APIResponse<T> {
  result?: T;
  error?: APIError;
}
