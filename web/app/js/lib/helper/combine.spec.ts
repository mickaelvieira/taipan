import combine from "./combine";

describe("combine helper", () => {
  test("creates an object with keys and values", () => {
    const keys = ["k1", "k2", "k3"];
    const values = ["v1", "v2", "v3"];
    const o = combine(keys, values);

    expect(Object.keys(o)).toEqual(keys);
    expect(o["k1"]).toBe("v1");
    expect(o["k2"]).toBe("v2");
    expect(o["k3"]).toBe("v3");
  });

  test("omits the last values if there are less keys than values", () => {
    const keys = ["k1", "k2"];
    const values = ["v1", "v2", "v3"];
    const o = combine(keys, values);

    expect(Object.keys(o)).toEqual(keys);
    expect(o["k1"]).toBe("v1");
    expect(o["k2"]).toBe("v2");
    expect(o["k3"]).toBeUndefined();
  });

  test("sets the values to undefined if there are less values than keys", () => {
    const keys = ["k1", "k2", "k3"];
    const values = ["v1", "v2"];
    const o = combine(keys, values);

    expect(Object.keys(o)).toEqual(keys);
    expect(o["k1"]).toBe("v1");
    expect(o["k2"]).toBe("v2");
    expect(o["k3"]).toBeUndefined();
  });
});
