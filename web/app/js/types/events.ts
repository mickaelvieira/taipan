export type BusTopic = "document" | "bookmark" | "user";

export type BusAction =
  | "add"
  | "remove"
  | "update"
  | "bookmark"
  | "unbookmark"
  | "favorite"
  | "unfavorite";

export interface Event {
  emitter: string;
  action: BusAction;
  topic: BusTopic;
}
