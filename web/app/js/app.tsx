import React from "react";
import { render } from "react-dom";
import { Provider } from "react-redux";
import "@babel/polyfill";
import App from "components/App";

import initFontAwesome from "services/fa";
import initMoment from "services/moment";
import initStore from "services/store";
import initDBStore from "services/idb";
import createSchema from "services/db/schema";
import { getDBState } from "services/repository";

import "../scss/styles.scss"

(function (win) {
  win.addEventListener("DOMContentLoaded", async () => {
    initFontAwesome();
    initMoment();
    initDBStore({ name: "bookmarks", version: 2 }, createSchema);

    console.time();
    const store = initStore(await getDBState());

    console.timeEnd();
    render(
      <Provider store={store}>
        <App />
      </Provider>,
      document.querySelector(".root")
    );
  });
})(window);
