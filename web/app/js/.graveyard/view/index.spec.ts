import { attachFragment, applyProperties, createWrapper } from "./index";

describe("attachFragment", () => {
  test("appends a node to another node", () => {
    const root = document.createElement("div");
    const element = document.createElement("div");

    attachFragment(root, element);

    expect(root.firstChild).toBe(element);
  });

  test("appends a node to another node but remove previous child nodes", () => {
    const root = document.createElement("div");
    root.innerHTML = "<div><p></p></div><div><p></p></div>";

    expect(root.childNodes.length).toBe(2);

    const element = document.createElement("div");

    attachFragment(root, element);

    expect(root.childNodes.length).toBe(1);
    expect(root.firstChild).toBe(element);
  });

  test("appends a fragment to node", () => {
    const root = document.createElement("div");
    const fragment = document.createDocumentFragment();
    const element = document.createElement("div");
    fragment.appendChild(element);

    attachFragment(root, fragment);

    expect(root.firstChild).toBe(element);
  });
});
