import React from "react";
import { render } from "react-dom";
import App from "./components/App";
import initSW from "./services/sw";

initSW(window.navigator);
render(React.createElement(App), document.querySelector(".root"));
