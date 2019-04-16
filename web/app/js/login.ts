import Login from "components/Login";
import "@babel/polyfill";

(function(win) {
  win.addEventListener("DOMContentLoaded", () => {
    Login();
  });
})(window);
