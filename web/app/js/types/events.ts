export type BusTopic = "News" | "Favorites" | "ReadingList" | "User";

export type BusAction = "Add" | "Remove" | "Update";

export interface Event {
  emitter: string;
  action: BusAction;
  topic: BusTopic;
}
