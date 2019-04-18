import React from "react";
import { render } from "react-dom";
import "@babel/polyfill";
import App from "./components/App";

render(React.createElement(App), document.querySelector(".root"));
