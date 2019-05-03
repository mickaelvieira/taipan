import React from "react";
import { render } from "react-dom";
import "@babel/polyfill";
import App from "./components/App";
import initMoment from "./services/moment";

initMoment();
render(React.createElement(App), document.querySelector(".root"));
