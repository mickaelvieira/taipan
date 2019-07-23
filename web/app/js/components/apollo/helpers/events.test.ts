import { isEmitter } from "./events";
import { Event } from "../../../types/events";

describe("Events helpers", () => {
  describe("isEmitter", () => {
    let event: Event;
    beforeAll(() => {
      event = {
        emitter: "foo",
        action: "favorite",
        topic: "bookmark"
      };
    });

    it("returns false if the event is null", () => {
      expect(isEmitter(null, "foo")).toBe(false);
    });

    it("returns false if the emitter does not match the client ID", () => {
      expect(isEmitter(event, "bar")).toBe(false);
    });

    it("returns false if the emitter match the client ID", () => {
      expect(isEmitter(event, "foo")).toBe(true);
    });
  });
});
