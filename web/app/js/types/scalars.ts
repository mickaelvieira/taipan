// Apollo client does not support custom scalar for now
// so this is just a workaround
export type Datetime = Date | string;
// export type Url = URL | string