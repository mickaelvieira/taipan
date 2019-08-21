import React from "react";
import { render } from "react-dom";
import App from "./components/app";
import initServiceWorker from "./services/sw";
import "regenerator-runtime/runtime";
import "emoji-mart/css/emoji-mart.css";

initServiceWorker(window);
render(React.createElement(App), document.querySelector(".root"));
