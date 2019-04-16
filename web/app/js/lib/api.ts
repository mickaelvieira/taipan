import { APP_DOMAIN_API } from "./constants";
import { CollectionResponse } from "collection/types";

export enum HTTPMethod {
  GET = "get",
  PUT = "put",
  POST = "post",
  DELETE = "delete"
}

export enum ContentTypes {
  JSON = "application/vnd.collection+json",
  HTML = "text/html",
  IMAGE = "*/*"
}

class Unauthorized extends Error {}

export default class API {

  static get(endpoint: string, contentType: ContentTypes = ContentTypes.JSON, options: Object | null = null) {
    return API.call(endpoint, HTTPMethod.GET, null, contentType, options);
  }

  static post(
    endpoint: string,
    data: Object | null = null,
    contentType: ContentTypes = ContentTypes.JSON, options: Object | null = null
  ) {
    return API.call(endpoint, HTTPMethod.POST, data, contentType, options);
  }

  static put(
    endpoint: string,
    data: Object | null = null,
    contentType: ContentTypes = ContentTypes.JSON, options: Object | null = null
  ) {
    return API.call(endpoint, HTTPMethod.PUT, data, contentType, options);
  }

  static delete(
    endpoint: string,
    contentType: ContentTypes = ContentTypes.JSON, options: Object | null = null
  ) {
    return API.call(endpoint, HTTPMethod.DELETE, null, contentType, options);
  }

  static isAuthorized(response: Response) {
    if (response.status === 401) {
      throw new Unauthorized();
    }

    return response;
  }

  static wasResponseSuccessful(response: Response): Response {
    if (!response.ok) {
      const msg = `Request to URL "${response.url}" has failed, status code "${
        response.status
      }"`;
      throw new Error(msg);
    }

    return response;
  }

  /**
   * Interrupt promises chain if the api does not return json
   */
  static isJsonResponse(response: Response): Response {
    const contentType = response.headers.get("content-type");
    const contentTypes = ["application/json", "application/vnd.collection+json"];

    if (!contentType || !contentTypes.includes(contentType)) {
      throw new Error("Response body does not appear to be of type JSON");
    }

    return response;
  }

  /**
   * Interrupt promises chain if the api returns an error
   */
  static isApiError(json: CollectionResponse): CollectionResponse {
    if ("error" in json.collection) {
      const error = JSON.stringify(json.collection.error);
      throw new Error(`The API has returned an error: ${error}`);
    }

    return json;
  }

  static call(
    endpoint: string,
    method: HTTPMethod,
    data: Object | null = null,
    contentType: ContentTypes = ContentTypes.JSON,
    opts: Object | null = null
  ): Promise<any> | [Promise<any>, Function] {
    const mode = "cors";
    const headers = new Headers();
    const encoding = "utf-8";
    const credentials = "include";
    const body = data ? JSON.stringify(data) : null;

    headers.append("Accept", contentType);
    headers.append("X-Requested-With", "XMLHttpRequest");

    const options: RequestInit = {
      ...opts,
      method,
      headers,
      credentials,
      mode
    };

    if (body) {
      Object.assign(options, { body });
      headers.append("Content-Type", contentType + "; charset=" + encoding);
    }

    const url = endpoint.startsWith("http")
      ? endpoint
      : `${APP_DOMAIN_API}${endpoint}`;
    const request = new Request(url, options);
    let promise = window.fetch(request);

    if (contentType !== ContentTypes.JSON) {
      return promise;
    }

    return promise
      .then(response => API.isAuthorized(response))
      .then(response => API.isJsonResponse(response))
      .then(response => response.json())
      .then(json => API.isApiError(json))
      .catch(function(error) {
        if (error instanceof Unauthorized) {
          window.location = "/login";
        }
      });
  }
}
