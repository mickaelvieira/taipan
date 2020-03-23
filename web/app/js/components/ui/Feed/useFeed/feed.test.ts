import {
  calculateCursor,
  calculateInterval,
  calculateHeightUpToFirstIndex,
  calculateHeightFromIndex,
} from "./helpers";

function interval(a: number, b: number): number {
  return b - a + 1;
}

describe("Feed helpers", () => {
  describe("calculateFirstIndex", () => {
    let heights: number[] = [];
    beforeAll(() => {
      heights = [10, 20, 10, 20, 10, 20, 10, 20, 10, 20];
    });
    it("calculates the gap up to the index", () => {
      [
        { g: 0, i: 0 },
        { g: 5, i: 0 },
        { g: 10, i: 1 },
        { g: 15, i: 1 },
        { g: 30, i: 2 },
        { g: 35, i: 2 },
        { g: 40, i: 3 },
        { g: 50, i: 3 },
        { g: 60, i: 4 },
        { g: 65, i: 4 },
        { g: 70, i: 5 },
        { g: 75, i: 5 },
        { g: 80, i: 5 },
        { g: 85, i: 5 },
        { g: 90, i: 6 },
        { g: 100, i: 7 },
        { g: 120, i: 8 },
        { g: 130, i: 9 },
        { g: 150, i: 10 },
      ].forEach(({ i, g }) => {
        expect(calculateCursor(g, heights)).toBe(i);
      });
    });
  });

  xdescribe("calculateBoundaries", () => {
    describe("when the total is lower than the page size", () => {
      it("selects all items, regardless the cursor's position", () => {
        for (let t = 0; t < 10; t++) {
          for (let c = 0; c < 9; c++) {
            const f = 0;
            const l = t === 0 ? 0 : t - 1;
            const [first, last] = calculateInterval(c, t);
            expect(first).toBe(f);
            expect(last).toBe(l);
          }
        }
      });
    });

    describe("when the total is up to a page and half", () => {
      it("selects all items, regardless the cursor's position", () => {
        for (let t = 0; t <= 15; t++) {
          for (let c = 0; c < 14; c++) {
            const f = 0;
            const l = t === 0 ? 0 : t - 1;
            const [first, last] = calculateInterval(c, t);
            expect(first).toBe(f);
            expect(last).toBe(l);
          }
        }
      });
    });

    describe("when the total is greater than a page and a half", () => {
      it("selects 10 items when the cursor is close to the beginning", () => {
        for (let t = 16; t <= 30; t++) {
          for (let c = 0; c <= 4; c++) {
            const f = 0;
            const l = 9;
            const [first, last] = calculateInterval(c, t);
            expect(first).toBe(f);
            expect(last).toBe(l);
          }
        }
      });

      it("selects 10 items when the cursor is close to the end", () => {
        for (let t = 16; t <= 30; t++) {
          for (let c = 25; c <= 29; c++) {
            const l = t - 1;
            const f = l - 9;
            const [first, last] = calculateInterval(c, t);
            expect(first).toBe(f);
            expect(last).toBe(l);
            expect(interval(first, last)).toBe(10);
          }
        }
      });

      it("moves the page with the cursor", () => {
        for (let t = 16; t <= 20; t++) {
          for (let c = 5; c <= 9; c++) {
            const f = c - 5;
            const l = c + 5;
            const [first, last] = calculateInterval(c, t);
            expect(first).toBe(f);
            expect(last).toBe(l);
            expect(interval(first, last)).toBe(11);
          }
        }
      });
    });
  });

  describe("calculateTopGapUptToFirstIndex", () => {
    let heights: number[] = [];
    beforeAll(() => {
      heights = [10, 20, 30, 40, 50, 60, 70, 80, 90, 100];
    });

    it("calculates the gap up to the index", () => {
      [
        { i: 0, g: 0 },
        { i: 1, g: 10 },
        { i: 2, g: 30 },
        { i: 3, g: 60 },
        { i: 4, g: 100 },
      ].forEach(({ i, g }) => {
        expect(calculateHeightUpToFirstIndex(i, heights)).toBe(g);
      });
    });
  });

  describe("calculateBottomGapFromLastIndex", () => {
    let heights: number[] = [];
    beforeAll(() => {
      heights = [10, 20, 30, 40, 50, 60, 70, 80, 90, 100];
    });

    it("calculates the gap from to the last index", () => {
      [
        { i: 5, g: 340 },
        { i: 6, g: 270 },
        { i: 7, g: 190 },
        { i: 8, g: 100 },
        { i: 9, g: 0 },
      ].forEach(({ i, g }) => {
        expect(calculateHeightFromIndex(i, heights)).toBe(g);
      });
    });
  });
});
