import React from "react";
import { render } from "react-dom";
import "@babel/polyfill";
import App from "./components/App";
import initSW from "./services/sw";
import initMoment from "./services/moment";

initSW(window.navigator);
initMoment();
render(React.createElement(App), document.querySelector(".root"));
