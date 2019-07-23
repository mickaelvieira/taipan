import { Event } from "../../../types/events";

export function isEmitter(event: Event | null, clientId: string): boolean {
  if (!event) {
    return false;
  }
  return event.emitter === clientId;
}
