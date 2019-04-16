import "@babel/polyfill";

customElements.define(
  "fancy-tabs",
  class extends HTMLElement {
    constructor() {
      super(); // always call super() first in the constructor.

      // Attach a shadow root to <fancy-tabs>.
      const shadowRoot = this.attachShadow({ mode: "open" });
      shadowRoot.innerHTML = `
      <div id="tabs">
        <slot id="tabsSlot" name="title"></slot>
      </div>
      <div id="panels">
        <slot id="panelsSlot"></slot>
      </div>
    `;
    }
  }
);

(function(win) {
  win.addEventListener("DOMContentLoaded", async () => {
    const header = document.createElement("header");
    const shadowRoot = header.attachShadow({ mode: "open" });
    shadowRoot.innerHTML =
      '<h1>Hello Shadow DOM</h1><div><slot name="cool"></slot></div>'; // Could also use appendChild().

    const content = document.createElement("div");
    content.setAttribute("slot", "cool");
    content.textContent = "Yo";
    header.appendChild(content);
    const body = document.querySelector("body");
    body.appendChild(header);
    console.log(body);
    console.log(customElements);
  });
})(window);
